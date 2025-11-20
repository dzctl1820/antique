package router

import (
	"an-backend/database"
	"an-backend/internal/handler"
	"an-backend/internal/service"
	"an-backend/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	database.InitGormDB()
	database.InitRedis() // 初始化Redis
	t := handler.UserService{DB: &service.UserDB{DB: database.DB}}
	v := r.Group("/api")
	{
		v.POST("/login", t.LoginByEmail)
		v.POST("/send/code", t.SendCodeByEmail)
		v.POST("/register", t.RegisterByEmail)
	}
	a := handler.ArticleService{DB: &service.ArticleDB{DB: database.DB}}
	p := r.Group("/api/article")
	p.Use(utils.AuthMiddleware())
	{
		p.POST("/add", a.AddArticle)
		p.GET("/one/:id", a.GetArticle)
		p.GET("/all", a.GetArticles)
	}
	c := handler.AiChatService{DB: &service.AiChatDB{DB: database.DB}}
	q := r.Group("/api/ai")
	q.Use(utils.AuthMiddleware())
	{
		q.POST("/chat", c.Chat)
		q.POST("/new/session", c.NewSession)
		q.GET("/sessions", c.GetSession)
		q.DELETE("/session", c.DeleteSession)
		q.POST("/session", c.UpdateSession)
		q.GET("/session/conversation", c.GetSessionConversation)
	}
	return r
}
