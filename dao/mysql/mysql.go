package mysql

import (
	"RedBubble/models"
	"RedBubble/setting"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// 声明一个全局的mdb变量
var mdb *gorm.DB

//使用GORM连接mysql
func Init(cfg *setting.MySQLConfig) (err error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//      username:password@protocol(address)/dbname?param=value
	//dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//1、打开数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	mdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //取消外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
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

	//2、配置数据库连接池
	sqlDB, err := mdb.DB() //获取通用数据库对象 sqlDB
	if err != nil {
		zap.L().Error("get *sql.DB failed", zap.Error(err))
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns) //10
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns) //20
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//3、创建数据库表，仅第一次的时候需要执行该方法
	//err = CreateSQLTable()
	//if err != nil {
	//	zap.L().Error("creat sql table failed", zap.Error(err))
	//	return
	//}

	return
}

//使用gorm而不是sql文件创建数据库表，在models中设置字段的规则如非空、添加索引，在这个方法里设置表的规则如字符集
func CreateSQLTable() (err error) {

	//user表
	err = mdb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment '用户表'").Migrator().CreateTable(&models.User{}) // 设置ENGINE=InnoDB，字符集=utf8mb4
	//category表
	err = mdb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment '帖子分类表'").Migrator().CreateTable(&models.Category{}) // 设置ENGINE=InnoDB，字符集=utf8mb4
	//post类
	err = mdb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment '帖子表'").Migrator().CreateTable(&models.Post{}) // 设置ENGINE=InnoDB，字符集=utf8mb4

	return
}
