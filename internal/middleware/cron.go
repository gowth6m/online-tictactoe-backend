package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-tictactoe/internal/api"
	"online-tictactoe/internal/config"
)

// CronJobAuthMiddleware checks if the Authorization header matches the CRON_SECRET set in Vercel
func CronJobAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cronJobSecret := config.AppConfig().Vercel.CronSecret
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
