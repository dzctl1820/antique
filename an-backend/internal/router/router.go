package router

import (
	"an-backend/database"
	"an-backend/internal/handler"
	"an-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	database.InitGormDB()
	t := handler.UserService{DB: &service.DataDB{DB: database.DB}}
	v := r.Group("/api")
	{
		v.POST("/login", t.LoginByEmail)
	}
	return r
}
