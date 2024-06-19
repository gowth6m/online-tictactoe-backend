package routes

import (
	"net/http"
	"online-tictactoe/internal/api"
	"online-tictactoe/internal/config"

	"github.com/gin-gonic/gin"
)

// For /
func DefaultRoutes(group *gin.Engine) {
	group.GET("/", func(c *gin.Context) {
		api.Success(c, http.StatusOK, "Online TicTacToe", gin.H{
			"message": "Welcome to the Online TicTacToe REST API v0",
			"version": config.AppConfig().App.AppVersion,
			"author":  "Gowthaman Ravindrathas",
		})
	})
}
