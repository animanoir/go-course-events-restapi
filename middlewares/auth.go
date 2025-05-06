package middlewares

import (
	"events-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	// Route protection.
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You're not authorized to do that, my baby.",
		})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
