package routes

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterForGame(c *gin.Context) {
	userId := c.GetInt64("uid")
	gameId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse id"})
		return
	}

	game, err := models.GetGame(gameId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find game"})
		return
	}

	err = game.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for game"})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for game"})
}

func CancelRegistration(c *gin.Context) {
	userId := c.GetInt64("uid")
	gameId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse id"})
		return
	}

	game, err := models.GetGame(gameId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find game"})
		return
	}

	err = game.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel registration"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully cancelled registration"})
}
