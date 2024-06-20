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
	privateGameGroup := group.Group("/game")
	privateGameGroup.Use(middleware.AuthMiddleware())
	cronJobGroup := group.Group("/game")
	cronJobGroup.Use(middleware.CronJobAuthMiddleware())

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
	privateGameGroup.POST("/:gameName/reset", func(c *gin.Context) {
		gameHandler.ResetGame(c)
	})

	// --- PUBLIC ROUTES ---
	gameGroup.GET("/all/count", func(c *gin.Context) {
		gameHandler.GetAllGamesCount(c)
	})

	// --- CRON JOB ROUTES ---
	cronJobGroup.POST("/all/clear", func(c *gin.Context) {
		gameHandler.ClearAllGames(c) // "0 5 */2 * *" => 05:00 AM, every 2 days
	})
}
