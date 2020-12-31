package bootstrap

import (
	c "gin_web/src/main/go/com.github.mushanwb.gin_web/Utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func SetupDB() {
	db, _ := ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}

func ConnectDB() (*gorm.DB, error) {
	var err error
	host := c.GetString("database.mysql.host")
	port := c.GetString("database.mysql.port")
	database := c.GetString("database.mysql.database")
	username := c.GetString("database.mysql.username")
	password := c.GetString("database.mysql.password")
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return DB, err
}
