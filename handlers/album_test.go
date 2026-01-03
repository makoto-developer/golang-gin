package handlers

import (
	"bytes"
	"encoding/json"
	"golang-gin/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

func TestGetAlbums(t *testing.T) {
	router := setupTestRouter()
	router.GET("/albums", GetAlbums)

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
	router := setupTestRouter()
	router.GET("/albums/:id", GetAlbumByID)

	tests := []struct {
		name           string
		id             string
		expectedStatus int
	}{
		{"Valid ID", "1", http.StatusOK},
		{"Invalid ID", "999", http.StatusNotFound},
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
				if album.ID != tt.id {
					t.Errorf("Expected album ID %s, got %s", tt.id, album.ID)
				}
			}
		})
	}
}

func TestPostAlbums(t *testing.T) {
	router := setupTestRouter()
	router.POST("/albums", PostAlbums)

	newAlbum := models.Album{
		ID:     "100",
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

	if createdAlbum.ID != newAlbum.ID {
		t.Errorf("Expected album ID %s, got %s", newAlbum.ID, createdAlbum.ID)
	}
	if createdAlbum.Title != newAlbum.Title {
		t.Errorf("Expected album title %s, got %s", newAlbum.Title, createdAlbum.Title)
	}
}

func TestPostAlbums_InvalidJSON(t *testing.T) {
	router := setupTestRouter()
	router.POST("/albums", PostAlbums)

	invalidJSON := []byte(`{"invalid": "json"`)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}
