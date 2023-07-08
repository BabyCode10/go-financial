package controllers

import (
	"fmt"
	"go-financial/app/config"
	"go-financial/app/models"
	"go-financial/app/requests"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type SAuthService struct {
	user models.IAuth
}

func Auth(user models.IAuth) *SAuthService {
	return &SAuthService{
		user: user,
	}
}

func (service SAuthService) Login(context *gin.Context) {
	var request requests.SAuthLoginRequest

	env, _ := config.Load()

	err := context.ShouldBindJSON(&request)

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	user, err := service.user.FindByUsername(request.Username)

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", user.Id),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(env.JWTKey))

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{
			"token": tokenString,
		},
	)
}

func (service SAuthService) Register(context *gin.Context) {
	var request requests.SAuthRegisterRequest

	fmt.Println(request)

	err := context.ShouldBindJSON(&request)

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	user, err := service.user.Store(&request)

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{
			"data": &user,
		},
	)
}

func (service SAuthService) Logout(context *gin.Context) {
	_, err := service.user.FindById(context.GetString("user_id"))

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"token": nil,
		},
	)
}

func (service SAuthService) Me(context *gin.Context) {
	user, err := service.user.FindById(context.GetString("user_id"))

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{
			"data": &user,
		},
	)
}
