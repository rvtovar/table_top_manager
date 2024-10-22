package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/games", GetGames)
	server.POST("/games", CreateGame)
	server.GET("/games/:id", GetGame)
	server.PUT("/games/:id", UpdateGame)
	server.DELETE("/games/:id", DeleteGame)
}
