package middleware

import (
	"net/http"
	"online-tictactoe/internal/api"
	"os"

	"github.com/gin-gonic/gin"
)

// CronJobAuthMiddleware checks if the Authorization header matches the CRON_SECRET set in Vercel
func CronJobAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cronJobSecret := os.Getenv("CRON_SECRET")
		if cronJobSecret == "" {
			api.Error(c, http.StatusInternalServerError, "CRON_SECRET not set", nil)
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader != "Bearer "+cronJobSecret {
			api.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
			return
		}

		c.Next()
	}
}
