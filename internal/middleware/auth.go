package middleware

import (
	"encoding/base64"
	"online-tictactoe/internal/api"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks for Basic auth token in the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			api.Error(c, 401, "Unauthorized", nil)
			return
		}

		// Basic authorization format: "Basic base64encoded(username:password)"
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Basic" {
			api.Error(c, 401, "Unauthorized", nil)
			return
		}

		authPayload, err := base64.StdEncoding.DecodeString(authParts[1])
		if err != nil {
			api.Error(c, 401, "Unauthorized", nil)
			return
		}

		credentials := strings.SplitN(string(authPayload), ":", 2)
		if len(credentials) != 2 {
			api.Error(c, 401, "Unauthorized", nil)
			return
		}

		userId := credentials[0]

		// Add userId to the context
		c.Set("userId", userId)

		// Continue to the next handler
		c.Next()
	}
}
