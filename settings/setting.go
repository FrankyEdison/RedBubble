package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//全局public变量Conf，用来保存所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	Version      string `mapstructure:"version"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"db_name"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
	Password string `mapstructure:"password"`
}

func Init() (err error) {
	//1、查找并读取配置文件
	viper.SetConfigFile("./config.yaml") // = SetConfigName + SetConfigType + AddConfigPath，这里使用的是相对路径：相对于可执行文件main.go的相对路径
	// 绝对路径：在磁盘中存放的绝对路径
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到
			fmt.Printf("未找到配置文件err : %v\n", err)
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Printf("其他err : %v\n", err)
		}
		return
	}

	//2、反序列viper对象并保存至全局变量Conf
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.unmarshal failed, err:%v \n", err)
	}

	//3、实时监控、读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) { //配置文件每次发生变更之后都会马上调用
		fmt.Println("Config file changed:", e.Name)
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.unmarshal failed, err:%v \n", err)
		}
	})

	return
}
