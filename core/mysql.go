package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"wTools/global"
)

func InitMysql() *gorm.DB {
	m := global.Config.Mysql
	dsn := m.DSN()

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库链接失败", m)
	}
	sdb, err := db.DB()
	sdb.SetMaxIdleConns(m.MaxIdleConns)
	sdb.SetMaxOpenConns(m.MaxOpenConns)
	sdb.SetConnMaxLifetime(10 * time.Minute)
	return db
}
