package main

import (
	"RedBubble/dao/mysql"
	"RedBubble/dao/redis"
	"RedBubble/logger"
	"RedBubble/router"
	"RedBubble/setting"
	"RedBubble/utils/snowflake"
	"RedBubble/utils/validator_"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//1、加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("加载配置失败, err:%v\n", err)
		return
	}

	//2、初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("初始化日志失败, err:%v\n", err)
		return
	}
	defer zap.L().Sync() //在程序退出前，把缓冲区的日志都写到磁盘中
	zap.L().Info("初始化日志成功")

	//3、初始化MySQL连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("初始化mysql连接失败, err:%v\n", err)
		return
	}

	//4、初始化Redis连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("初始化redis连接失败, err:d%v\n", err)
		return
	}
	defer redis.Close()

	//5、初始化雪花算法库
	if err := snowflake.Init(setting.Conf.SnowFlakeConfig.StartTime, setting.Conf.SnowFlakeConfig.MachineID); err != nil {
		fmt.Printf("初始化雪花算法库失败, err:%v\n", err)
		return
	}

	//6、初始化validator的翻译器，param="zh"即要翻译成中文
	if err := validator_.InitTrans("zh"); err != nil {
		fmt.Printf("初始化validator翻译器失败, err:%v\n", err)
		return
	}

	//7、注册路由
	r := router.Setup()

	//8、启动服务（优雅关机）
	//上面代码虽然用到gin来创建路由，但是实际上最终还是用go/http的方法来自定义server和启动server，因为优雅关机是 http.Server内置的方法
	//8.1、将路由引擎看做handler，自定义一个server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.Port),
		Handler: r,
	}

	//8.2、开启一个goroutine启动server，为什么要开协程去执行？因为ListenAndServe执行时会不停循环监听请求路径，导致程序无法往下执行
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: ", zap.Error(err))
		}
	}()

	//8.3、优雅关闭server，等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道，缓冲区大小为1
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以输入他没用
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 当quit通道一直没值，就会一直阻塞，直到接收到上述两种信号时才会往下执行
	zap.L().Info("关闭服务器 ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("服务器成功退出")

}
