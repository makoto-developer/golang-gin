package handlers

import (
	"bytes"
	"encoding/json"
	"golang-gin/models"
	"golang-gin/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() (*gin.Engine, *AlbumHandler) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockRepo := repository.NewMockAlbumRepository()
	handler := NewAlbumHandler(mockRepo)
	return router, handler
}

func TestGetAlbums(t *testing.T) {
	router, handler := setupTestRouter()
	router.GET("/albums", handler.GetAlbums)

	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var albums []models.Album
	err := json.Unmarshal(w.Body.Bytes(), &albums)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(albums) == 0 {
		t.Error("Expected at least one album")
	}

	t.Logf("Got %d albums", len(albums))
}

func TestGetAlbumByID(t *testing.T) {
	router, handler := setupTestRouter()
	router.GET("/albums/:id", handler.GetAlbumByID)

	tests := []struct {
		name           string
		id             string
		expectedStatus int
	}{
		{"Valid ID", "1", http.StatusOK},
		{"Invalid ID", "999", http.StatusNotFound},
		{"Invalid ID format", "abc", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/albums/"+tt.id, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var album models.Album
				err := json.Unmarshal(w.Body.Bytes(), &album)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}
				if album.ID == 0 {
					t.Error("Expected valid album ID")
				}
			}
		})
	}
}

func TestPostAlbums(t *testing.T) {
	router, handler := setupTestRouter()
	router.POST("/albums", handler.PostAlbums)

	newAlbum := models.Album{
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  19.99,
		Tax:    0.1,
	}

	jsonData, _ := json.Marshal(newAlbum)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var createdAlbum models.Album
	err := json.Unmarshal(w.Body.Bytes(), &createdAlbum)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if createdAlbum.ID == 0 {
		t.Error("Expected valid album ID")
	}
	if createdAlbum.Title != newAlbum.Title {
		t.Errorf("Expected album title %s, got %s", newAlbum.Title, createdAlbum.Title)
	}
}

func TestPostAlbums_InvalidJSON(t *testing.T) {
	router, handler := setupTestRouter()
	router.POST("/albums", handler.PostAlbums)

	invalidJSON := []byte(`{"invalid": "json"`)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}
