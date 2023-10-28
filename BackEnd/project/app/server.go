package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"project/app/controller"
	"project/source/app/services"
	"project/source/domain/entity"
	"project/source/infrastructure/postgresql"
	"project/source/infrastructure/repositories"
	pb "project/your/go/package" // Replace with the path to your proto package
)

var (
	//productService repositories.Repository = repositories.New()

	cfg                   = entity.GetConfig()
	postgreSQLClient, err = postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	repository            = repositories.New(postgreSQLClient)
	service               = services.NewService(&repository)
	Controller            = controller.New(*service)
)

const (
	address = "localhost:50051"
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
	gRPC()
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
func gRPC() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewYourServiceClient(conn)

	login := "YourTestLogin"
	password := "YourTestPassword"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	log.Printf("Server listening at %v", address)

	resp, err := client.Login(ctx, &pb.YourLoginRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not login: %v", err)
	}

	log.Printf("Received token: %s", resp)
}
