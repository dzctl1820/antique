package main

import (
	"an-backend/database"
	"an-backend/internal/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// 加载 .env 文件
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitGormDB()
	routers := router.InitRouter()
	err := routers.Run(":8086")
	if err != nil {
		return
	}
}
