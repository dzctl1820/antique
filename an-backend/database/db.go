package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGormDB() {
	dbMsg := "root:tiange1820@tcp(127.0.0.1:3306)/antique?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dbMsg), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	} else {
		println("数据库连接成功")
	}
}
