package routes

import (
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/games", GetGames)

	server.GET("/games/:id", GetGame)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/games", CreateGame)
	authenticated.PUT("/games/:id", UpdateGame)
	authenticated.DELETE("/games/:id", DeleteGame)

	authenticated.POST("/games/:id/register", RegisterForGame)
	authenticated.DELETE("/games/:id/register", CancelRegistration)

	// User Routes
	server.POST("/signup", SignUp)
	server.POST("/login", Login)
}
