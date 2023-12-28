package controller

import (
	_ "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"project/source/app/services"
	"project/source/infrastructure/utils"
	pb "project/your/go/package"
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
	Username         string `json:"username"`
	UserEmail        string `json:"user_email"`
	Password         string `json:"user_password"`
	ActivationStatus string `json:"user_activation_status"`
	Status           int    `json:"status"`
}

type tokenBody struct {
	Token string `json:"token"`
}

type Controller struct {
	Service    services.Service
	GRPCClient pb.YourServiceClient
}

func New(service services.Service, grpcClient pb.YourServiceClient) Controller {
	return Controller{
		Service:    service,
		GRPCClient: grpcClient,
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
	if !regexp.MustCompile(`^\S+@\S+\.\S+$`).MatchString(body.UserEmail) {
		return fmt.Errorf("invalid email format")
	}

	// Проверка пароля
	if len(body.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}

	// Проверка статуса активации
	if body.ActivationStatus != "active" && body.ActivationStatus != "inactive" {
		return fmt.Errorf("invalid activation status")
	}

	// Проверка статуса (можно добавить конкретные условия в зависимости от вашего приложения)

	return nil
}

func (c *Controller) SignUp(ctx *gin.Context) {
	var body SignUpBody
	if err := ctx.BindJSON(&body); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}
	fmt.Println(ctx.Request.Body)
	if err := ValidateSignUpBody(body); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response, err := c.GRPCClient.Register(ctx, &pb.YourRegistrationRequest{
		Login:    body.Username,
		Password: body.Password,
	})

	fmt.Println(response)
	//hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	//if err != nil {
	//	utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//user := entity.User{0, body.Username, body.UserEmail,
	//	string(hash), body.ActivationStatus, body.Status}
	//err = c.Service.SignUpService(user)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "created",
	})
}

func (c *Controller) Login(ctx *gin.Context) {
	err := ctx.BindJSON(&body)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	//user := entity.User{0, body.Username, body.UserEmail,
	//	body.Password, body.ActivationStatus, body.Status}
	response, err := c.GRPCClient.Login(ctx, &pb.YourLoginRequest{
		Login:    body.Username,
		Password: body.Password,
	})
	fmt.Println(response)
	//user, err = c.Service.FindUserId(user)
	//
	//if err != nil {
	//	utils.NewErrorResponse(ctx, http.StatusNotFound, "User not found or incorrect credentials")
	//	return
	//}
	//user, err = c.Service.FindUserPass(user)
	//if err != nil {
	//	utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	//	return
	//}
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"sub": user.UserID,
	//	"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	//})
	//tokenString, err := token.SignedString([]byte("sdff32dsadsadsdsds34sr2134rewtFSFSFSFASFASFASFASFASFASFASF3t2sra"))
	//if err != nil {
	//	utils.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//ctx.SetSameSite(http.SameSiteLaxMode)
	//ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	//
	//ctx.JSON(http.StatusOK, map[string]interface{}{
	//	"token": tokenString,
	//})
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
