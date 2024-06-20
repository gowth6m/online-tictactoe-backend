package serverless

import (
	"log"
	"online-tictactoe/internal/config"
	"online-tictactoe/internal/db"
	"online-tictactoe/internal/pusher"
	"online-tictactoe/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// This function initializes the serverless application.
// It loads the configuration, connects to the MongoDB database, sets up the routes, and returns the router.
//
// This is done to expose internal functions to the serverless environment.
//
// Returns: router, a clean up function to disconnect from the MongoDB database
func Initialize() (*gin.Engine, func()) {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	db.ConnectToMongoDB()
	pusher.Init()

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))
	return router, func() {
		db.DisconnectFromMongoDB()
	}
}
