package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"project/app/controller"
	"project/source/app/services"
	"project/source/domain/entity"
	"project/source/infrastructure/postgresql"
	"project/source/infrastructure/repositories"
)

var (
	//productService repositories.Repository = repositories.New()

	cfg                   = entity.GetConfig()
	postgreSQLClient, err = postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	repository            = repositories.New(postgreSQLClient)
	service               = services.NewService(&repository)
	Controller            = controller.New(*service)
)

func main() {
	server := gin.Default()
	//server.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{"http://localhost:4200"},
	//}))
	server.Use(CORSMiddleware())
	user := server.Group("/user")
	{
		user.POST("/register", Controller.SignUp)
		user.POST("/login", Controller.Login)
		user.POST("/", Controller.UserId)
		user.GET("/", Controller.Validate)
	}
	events := server.Group("/event")
	{
		//events.GET("/:id", Controller.Validate, Controller.FindUserEvents)
		events.GET("/:id", Controller.FindUserEvents)
	}
	server.Run()

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
