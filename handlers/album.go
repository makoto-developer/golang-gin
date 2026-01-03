package handlers

import (
	"golang-gin/models"
	"golang-gin/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AlbumHandler handles album-related requests
type AlbumHandler struct {
	repo repository.AlbumRepository
}

// NewAlbumHandler creates a new AlbumHandler
func NewAlbumHandler(repo repository.AlbumRepository) *AlbumHandler {
	return &AlbumHandler{repo: repo}
}

// GetAlbums returns all albums
func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	albums, err := h.repo.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch albums"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumByID returns a specific album by ID
func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	album, err := h.repo.FindByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch album"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

// PostAlbums adds a new album
func (h *AlbumHandler) PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
