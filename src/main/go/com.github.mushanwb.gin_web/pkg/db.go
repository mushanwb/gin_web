package pkg

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DefaultDB *gorm.DB

func ConnectDB(mysqlConfig map[string]interface{}) (*gorm.DB, error) {
	var err error
	host, _ := mysqlConfig["host"].(string)
	port, _ := mysqlConfig["port"].(string)
	database, _ := mysqlConfig["database"].(string)
	username, _ := mysqlConfig["username"].(string)
	password, _ := mysqlConfig["password"].(string)

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Logger.Error(err)
	}

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db, err
}

func ConnectDefaultDB(defaultConfig map[string]interface{}) {
	DefaultDB, _ = ConnectDB(defaultConfig)
}
