package middlewares

import (
	"errors"
	"net/http"
	"project-mygram/database"
	"project-mygram/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.SetupDbConnection()
		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Unauthorized",
				"error":   "Invalid ID data type",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint64(userData["id"].(float64))
		comment := entity.Comment{}
		user := entity.User{}
		err = db.First(&user, userID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Unauthorized",
				"error":   "Failed to find user",
			})
			return
		}

		err = db.Select("user_id").First(&comment, uint64(ID)).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Data not found",
				"error":   err.Error(),
			})
			return
		}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Unauthorized",
				"error":   "Failed to find social media",
			})
			return
		}

		if comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "Forbidden",
				"error":   "You are not allowed to access this comment",
			})
			return
		}

		c.Next()
	}
}
