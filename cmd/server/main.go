package main

import (
	"log"
	"online-tictactoe/internal/config"
	"online-tictactoe/internal/db"
	"online-tictactoe/internal/pusher"
	"online-tictactoe/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Online TicTacToe API
// @version 1
// @description This is the REST API for Online TicTacToe.
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BasicAuth
// @in header
// @name Authorization
func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	db.ConnectToMongoDB()
	pusher.Init()
	defer db.DisconnectFromMongoDB()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))
	routes.SetupRoutes(router)
	router.Run(config.AppConfig().App.Host + ":" + config.AppConfig().App.Port)
}
