package middlewares

import (
	"awesomeProject1/intelnal/apperror"
	"awesomeProject1/intelnal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(apperror.Unauthorized("Authorization header is missing", nil))
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) != 2 || fields[0] != "Bearer" {
			c.Error(apperror.Unauthorized("Invalid Authorization header", nil))
			return
		}
		tokenString := fields[1]
		claims, err := utils.ValidateAccessToken(tokenString)
		if err != nil {
			c.Error(apperror.Unauthorized("Invalid access token", err))
			return
		}

		if uid, ok := claims["user_id"].(float64); ok {
			c.Set("user_id", uint(uid))
		} else {
			c.Error(apperror.Unauthorized("Invalid user_id in token", nil))
			return
		}
		c.Next()
	}
}
