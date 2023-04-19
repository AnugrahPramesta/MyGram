package helpers

import (
	"errors"
	"fmt"
	"project-mygram/configs"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(id uint64, username, email string) (token string, err error) {
	configs := configs.GetInstance()
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = parseToken.SignedString([]byte(configs.Appconfig.SecretKey))

	return
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	configs := configs.GetInstance()
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errors.New("bearer token not found")
	}

	stringToken := headerToken[7:]
	fmt.Println(stringToken)

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("failed to get sign token")
		}

		return []byte(configs.Appconfig.SecretKey), nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, errors.New("failed to parse claims")
	}

	return token.Claims.(jwt.MapClaims), nil
}
