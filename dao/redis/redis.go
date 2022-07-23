package redis

import (
	"RedBubble/settings"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

// 声明一个全局的rdb变量，是private的
var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		//连的虚拟机的redis7
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize, // 连接池大小
	})
	//用于对连接数据库进行超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result() //判断是否连接成功，如果5秒都连接不上redis，就说明连接超时连接失败
	//defer rdb.Close() // 这行defer一般不在这里写，一般在main函数的最后写
	return err
}

//不暴露rdb变量，只暴露方法
func Close() {
	_ = rdb.Close()
}
