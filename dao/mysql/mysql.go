package mysql

import (
	"RedBubble/settings"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 声明一个全局的mdb变量
var mdb *gorm.DB

//使用GORM连接mysql
func Init(cfg *settings.MySQLConfig) (err error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//      username:password@protocol(address)/dbname?param=value
	//dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	mdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("open mysql failed", zap.Error(err))
		//fmt.Printf("open mysql failed, err:%v\n", err)
		return
	}
	//也可进行高级配置
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	//	DefaultStringSize: 256, // string 类型字段的默认长度
	//	DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})

	sqlDB, err := mdb.DB() //获取通用数据库对象 sqlDB
	if err != nil {
		zap.L().Error("get *sql.DB failed", zap.Error(err))
		//fmt.Printf("get *sql.DB failed, err:%v\n", err)
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns) //10

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns) //100

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return
}
