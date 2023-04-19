package middlewares

import (
	"fmt"
	"net/http"
	"project-mygram/helpers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := helpers.VerifyToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}

		ctx.Set("userData", claims)
		ctx.Next()
	}
}

func CheckAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		fmt.Println(userData)
		userID := uint64(userData["id"].(float64))
		if userID == 0 {
			return
		}

		ctx.Next()
	}
}
