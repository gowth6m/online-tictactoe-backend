package routes

import (
	"github.com/gin-gonic/gin"
	"online-tictactoe/internal/middleware"
	"online-tictactoe/internal/services/game"
)

// GROUP: /game
func GameRoutes(group *gin.Engine) {
	gameRepo := game.NewGameRepository()
	gameHandler := game.NewGameHandler(gameRepo)

	gameGroup := group.Group("/game")
	privateGameGroup := gameGroup.Use(middleware.AuthMiddleware())

	// --- PRIVATE ROUTES ---
	privateGameGroup.POST("/create", func(c *gin.Context) {
		gameHandler.CreateGame(c)
	})
	privateGameGroup.GET("/:gameName", func(c *gin.Context) {
		gameHandler.GetGame(c)
	})
	privateGameGroup.POST("/move", func(c *gin.Context) {
		gameHandler.CreateMove(c)
	})

	// --- PUBLIC ROUTES ---
	gameGroup.GET("/all/count", func(c *gin.Context) {
		gameHandler.GetAllGamesCount(c)
	})
}
