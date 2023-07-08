package middlewares

import (
	"fmt"
	"go-financial/app/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		env, _ := config.Load()

		accessToken := RequestHeader(context)

		if accessToken == "" {
			context.JSON(
				http.StatusUnauthorized,
				gin.H{"message": "Unauthorized."},
			)

			context.Abort()

			return
		}

		keyFunc := func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte([]byte(env.JWTKey)), nil
		}

		token, err := jwt.Parse(accessToken, keyFunc)

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

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": "Token not found.",
				},
			)

			context.Abort()

			return
		}

		sub := fmt.Sprint(claims["sub"])

		context.Set("user_id", sub)

		context.Next()
	}
}

func RequestHeader(context *gin.Context) string {
	var accessToken string

	bearerToken := context.Request.Header.Get("Authorization")

	fields := strings.Fields(bearerToken)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	}

	return accessToken
}
