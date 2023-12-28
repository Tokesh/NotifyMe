package controller

import (
	_ "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"log"
	"net/http"
	pb "project/proto"
	"project/source/app/services"
	"project/source/domain/entity"
	"project/source/infrastructure/utils"
	"regexp"
	"time"
)

var body struct {
	Username         string `json:"username"`
	UserEmail        string `json:"user_email"`
	Password         string `json:"user_password"`
	ActivationStatus string `json:"user_activation_status"`
	Status           int    `json:"status"`
}

type SignUpBody struct {
	Username         string
	UserEmail        string
	Password         string
	ActivationStatus string
	Status           int
}

type tokenBody struct {
	Token string `json:"token"`
}

type Controller struct {
	Service services.Service
}

func New(service services.Service) Controller {
	return Controller{
		Service: service,
	}
}

func ValidateSignUpBody(body SignUpBody) error {
	// Проверка имени пользователя
	if len(body.Username) < 3 || len(body.Username) > 30 {
		return fmt.Errorf("username must be between 3 and 30 characters")
	}
	if !regexp.MustCompile(`^\w+$`).MatchString(body.Username) {
		return fmt.Errorf("username must be alphanumeric")
	}

	// Проверка адреса электронной почты
	//if !regexp.MustCompile(`^\S+@\S+\.\S+$`).MatchString(body.UserEmail) {
	//	return fmt.Errorf("invalid email format")
	//}

	// Проверка пароля
	//if len(body.Password) < 6 {
	//	return fmt.Errorf("password must be at least 6 characters long")
	//}

	// Проверка статуса активации
	//if body.ActivationStatus != "active" && body.ActivationStatus != "inactive" {
	//	return fmt.Errorf("invalid activation status")
	//}

	// Проверка статуса (можно добавить конкретные условия в зависимости от вашего приложения)

	return nil
}

func (c *Controller) SignUp(ctx *gin.Context) {
	var body SignUpBody
	if err := ctx.BindJSON(&body); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := ValidateSignUpBody(body); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()
	client := pb.NewYourServiceClient(conn)
	fmt.Println(body.Password)
	fmt.Println(body.Username)
	// Выполнение gRPC запроса на регистрацию
	response, err := client.Register(ctx, &pb.YourRegistrationRequest{
		Login:    body.Username,
		Password: body.Password,
		// Добавьте другие необходимые поля
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": response.GetMessage(),
	})
}

func (c *Controller) Login(ctx *gin.Context) {
	err := ctx.BindJSON(&body)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		fmt.Printf("2")
		return
	}

	user := entity.User{0, body.Username, body.UserEmail,
		body.Password, body.ActivationStatus, body.Status}
	user, _ = c.Service.FindUserId(user)

	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()
	client := pb.NewYourServiceClient(conn)

	// Выполнение gRPC запроса
	fmt.Println(body.Username)
	fmt.Println(body.Password)
	response, err := client.Login(ctx, &pb.YourLoginRequest{
		Login:    body.Username, // Замените на ваше имя пользователя
		Password: body.Password, // Замените на ваш пароль
	})
	if err != nil {
		//log.Fatalf("Ошибка при выполнении запроса: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"token": err.Error(),
		})

	}
	log.Printf("Ответ от сервера: %s", response.GetMessage())

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": response.GetMessage(),
	})
}
func (c *Controller) Validate(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("sdff3234sr2134rewt3t2sra"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		user, err := c.Service.FindUserByIdService(int(claims["sub"].(float64)))
		if user.UserID == 0 {
			utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	} else {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": user,
	})
}

func (c *Controller) UserId(ctx *gin.Context) {
	tokeni := tokenBody{}
	err := ctx.BindJSON(&tokeni)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	tokenString := tokeni.Token
	fmt.Println(tokeni)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("sdff3234sr2134rewt3t2sra"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		user, err := c.Service.FindUserByIdService(int(claims["sub"].(float64)))
		if user.UserID == 0 {
			utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"user_id": user.UserID,
		})
	} else {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

}
