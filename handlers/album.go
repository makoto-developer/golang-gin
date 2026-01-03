package handlers

import (
	"golang-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAlbums returns all albums
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

// GetAlbumByID returns a specific album by ID
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// PostAlbums adds a new album
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
