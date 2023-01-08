package controller

import (
	_ "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"project/source/app/services"
	"project/source/domain/entity"
	"project/source/infrastructure/utils"
	"time"
)

var body struct {
	Username         string `json:"username"`
	UserEmail        string `json:"user_email"`
	Password         string `json:"user_password"`
	ActivationStatus string `json:"user_activation_status"`
	Status           int    `json:"status"`
}

type Controller struct {
	Service services.Service
}

func New(service services.Service) Controller {
	return Controller{
		Service: service,
	}
}

func (c *Controller) SignUp(ctx *gin.Context) {
	err := ctx.BindJSON(&body)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := entity.User{0, body.Username, body.UserEmail,
		string(hash), body.ActivationStatus, body.Status}
	err = c.Service.SignUpService(user)
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

	user := entity.User{0, body.Username, body.UserEmail,
		body.Password, body.ActivationStatus, body.Status}
	user = c.Service.FindUserId(user)

	if user.UserID == 0 {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user = c.Service.FindUserPass(user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("sdff3234sr2134rewt3t2sra"))
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": tokenString,
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
