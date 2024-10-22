package routes

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetGame(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse id"})
		return
	}
	game, err := models.GetGame(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get game"})
		return
	}
	c.JSON(http.StatusOK, game)
}
func GetGames(c *gin.Context) {
	games, err := models.AllGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, games)
}

func CreateGame(c *gin.Context) {

	var game models.Game
	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid := c.GetInt64("uid")
	game.UserId = uid
	game.Save()
	c.JSON(http.StatusCreated, game)
}

func UpdateGame(c *gin.Context) {
	gid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	uid := c.GetInt64("uid")
	game, err := models.GetGame(gid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find game"})
		return
	}

	if uid != game.UserId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	var updatedGame models.Game
	err = c.ShouldBindJSON(&updatedGame)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request"})
	}

	updatedGame.ID = gid
	err = updatedGame.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updatedGame)
}

func DeleteGame(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse id"})
		return
	}
	uid := c.GetInt64("uid")
	game, err := models.GetGame(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find game"})
		return
	}

	if uid != game.UserId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}
	err = game.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete game"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Game deleted"})
}
