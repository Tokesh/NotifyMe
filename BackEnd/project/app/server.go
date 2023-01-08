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
	user := server.Group("/user")
	{
		user.POST("/register", Controller.SignUp)
		user.POST("/login", Controller.Login)
		user.GET("/", Controller.Validate)
	}
	events := server.Group("/event")
	{
		events.GET("/:id", Controller.Validate, Controller.FindUserEvents)
	}
	server.Run()

}
