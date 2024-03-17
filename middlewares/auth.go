package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"test.com/event-api/utils"
)

func Authenticate(context *gin.Context) {

	// token := context.Request.Header.Get("Authorization")
	token := context.GetHeader("Authorization")

	// token = strings.Split(token, "Bearer ")[1]

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not unauthorized",
		})
		return
	}

	tokenData, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Not Authorized",
			"error":   err.Error(),
		})
		return
	}
	expireTime := int64(tokenData["exp"].(float64))
	if expireTime < time.Now().Unix() {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "jwt token Expired",
		})
		return
	}

	userId := int64(tokenData["userId"].(float64))

	context.Set("userId", userId)
	context.Next()
}
