package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Swagger routes for API documentation
	SwaggerRoutes(router)

	// Default routes for the root path
	DefaultRoutes(router)

	// Game routes for the /game path
	GameRoutes(router)
}
