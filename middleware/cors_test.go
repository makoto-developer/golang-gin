package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCORS(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(CORS())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	t.Run("GET request with CORS headers", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		// Check CORS headers
		expectedHeaders := map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		}

		for header, expected := range expectedHeaders {
			actual := w.Header().Get(header)
			if actual != expected {
				t.Errorf("Expected header %s to be '%s', got '%s'", header, expected, actual)
			}
		}
	})

	t.Run("OPTIONS preflight request", func(t *testing.T) {
		req, _ := http.NewRequest("OPTIONS", "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusNoContent {
			t.Errorf("Expected status %d, got %d", http.StatusNoContent, w.Code)
		}

		// Check CORS headers exist
		if w.Header().Get("Access-Control-Allow-Origin") == "" {
			t.Error("Expected Access-Control-Allow-Origin header")
		}
		if w.Header().Get("Access-Control-Allow-Methods") == "" {
			t.Error("Expected Access-Control-Allow-Methods header")
		}
	})
}
