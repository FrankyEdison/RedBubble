package logger

import (
	"RedBubble/settings"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/**
	1、自定义logger，自定义的内容包括日志保存到哪里、日志应该含什么内容、日志的输出格式等。并且将自定义logger替换掉zap包中全局的logger实例
	2和3、自定义GinLogger和GinRecovery中间件，GinLogger和GinRecovery是重写了gin默认的logger和recovery中间件中的日志相关方法，
          至于重写的内容就是用我们自定义的logger代替原来默认的日志相关方法
*/

//1、初始化Logger
func Init(cfg *settings.LogConfig) (err error) {
	//1.1 设置WriterSyncer
	writeSyncer := getLogWriter(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
	)
	//1.2 设置encoder
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	//1.3 创建自定义core
	core := zapcore.NewCore(encoder, writeSyncer, l)
	//1.4 创建自定义logger
	lg := zap.New(core, zap.AddCaller()) //参数zap.AddCaller()的作用是在日志中记录go程序的哪一行调用本日志
	//1.5 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(lg)
	return
}

//1.1设置WriterSyncer
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  //**日志文件的位置和名字
		MaxSize:    maxSize,   //**在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: maxBackup, //**保留旧文件的最大个数
		MaxAge:     maxAge,    //**保留旧文件的最大天数
		Compress:   false,     //**是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

//1.2设置encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //**修改为可读的日期格式
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //**使用大写字母记录日志级别
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

//2、自定义GinLogger中间件 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		//lg的地方就是自定义的地方
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

//3、 自定义GinRecovery中间件 recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					//zap.L()的地方就是自定义的地方
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					//zap.L()的地方就是自定义的地方
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					//zap.L()的地方就是自定义的地方
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
