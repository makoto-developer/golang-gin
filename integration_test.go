package main

import (
	"bytes"
	"encoding/json"
	"golang-gin/handlers"
	"golang-gin/middleware"
	"golang-gin/models"
	"golang-gin/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupIntegrationTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(gin.Recovery())

	// Initialize mock repository and handler
	mockRepo := repository.NewMockAlbumRepository()
	albumHandler := handlers.NewAlbumHandler(mockRepo)

	// Health check
	router.GET("/health", handlers.HealthCheck)

	// Album routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/albums", albumHandler.GetAlbums)
		v1.GET("/albums/:id", albumHandler.GetAlbumByID)
		v1.POST("/albums", albumHandler.PostAlbums)
	}

	return router
}

func TestIntegration_FullWorkflow(t *testing.T) {
	router := setupIntegrationTestRouter()

	t.Run("1. Health Check", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Health check failed: status %d", w.Code)
		}

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		if response["status"] != "healthy" {
			t.Error("Health check status not healthy")
		}
	})

	t.Run("2. Get All Albums", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/albums", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Get albums failed: status %d", w.Code)
		}

		var albums []models.Album
		json.Unmarshal(w.Body.Bytes(), &albums)
		if len(albums) == 0 {
			t.Error("Expected at least one album")
		}
	})

	t.Run("3. Get Album by ID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/albums/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Get album by ID failed: status %d", w.Code)
		}

		var album models.Album
		json.Unmarshal(w.Body.Bytes(), &album)
		if album.ID != 1 {
			t.Errorf("Expected album ID 1, got %d", album.ID)
		}
	})

	t.Run("4. Create New Album", func(t *testing.T) {
		newAlbum := models.Album{
			Title:  "Integration Test Album",
			Artist: "Test Artist",
			Price:  29.99,
			Tax:    0.1,
		}

		jsonData, _ := json.Marshal(newAlbum)
		req, _ := http.NewRequest("POST", "/api/v1/albums", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("Create album failed: status %d", w.Code)
		}

		var createdAlbum models.Album
		json.Unmarshal(w.Body.Bytes(), &createdAlbum)
		if createdAlbum.ID == 0 {
			t.Error("Expected non-zero album ID")
		}
	})

	t.Run("5. CORS Headers", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/albums", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Header().Get("Access-Control-Allow-Origin") != "*" {
			t.Error("CORS header not set correctly")
		}
	})

	t.Run("6. 404 Not Found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/albums/999", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", w.Code)
		}
	})

	t.Run("7. Invalid ID format", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/albums/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}
	})
}

func TestIntegration_ErrorHandling(t *testing.T) {
	router := setupIntegrationTestRouter()

	t.Run("Invalid JSON", func(t *testing.T) {
		invalidJSON := []byte(`{"invalid": json}`)
		req, _ := http.NewRequest("POST", "/api/v1/albums", bytes.NewBuffer(invalidJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}
	})

	t.Run("Missing Content-Type", func(t *testing.T) {
		newAlbum := models.Album{Title: "Test", Artist: "Test", Price: 10, Tax: 0.1}
		jsonData, _ := json.Marshal(newAlbum)
		req, _ := http.NewRequest("POST", "/api/v1/albums", bytes.NewBuffer(jsonData))
		// Not setting Content-Type header
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Should still work with Gin's auto-detection
		if w.Code != http.StatusCreated && w.Code != http.StatusBadRequest {
			t.Logf("Status: %d (acceptable)", w.Code)
		}
	})
}
