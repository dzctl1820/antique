package main

import (
	"an-backend/database"
	"an-backend/internal/router"
)

func main() {
	database.InitGormDB()
	routers := router.InitRouter()
	err := routers.Run(":8086")
	if err != nil {
		return
	}
}
