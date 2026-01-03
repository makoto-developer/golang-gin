package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MockUser represents a mock user response
type MockUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
}

// MockProduct represents a mock product response
type MockProduct struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "http-mock-server",
		})
	})

	// Mock API endpoints
	api := router.Group("/api/v1")
	{
		// Get users
		api.GET("/users", func(c *gin.Context) {
			users := []MockUser{
				{ID: 1, Name: "Alice", Email: "alice@example.com", UserType: "admin"},
				{ID: 2, Name: "Bob", Email: "bob@example.com", UserType: "user"},
				{ID: 3, Name: "Charlie", Email: "charlie@example.com", UserType: "user"},
			}
			c.JSON(http.StatusOK, users)
		})

		// Get user by ID
		api.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			user := MockUser{
				ID:       1,
				Name:     "Alice",
				Email:    "alice@example.com",
				UserType: "admin",
			}
			log.Printf("Requested user ID: %s", id)
			c.JSON(http.StatusOK, user)
		})

		// Create user
		api.POST("/users", func(c *gin.Context) {
			var user MockUser
			if err := c.BindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			user.ID = 999
			log.Printf("Created user: %+v", user)
			c.JSON(http.StatusCreated, user)
		})

		// Get products
		api.GET("/products", func(c *gin.Context) {
			products := []MockProduct{
				{ID: 1, Name: "Laptop", Price: 1200.00},
				{ID: 2, Name: "Mouse", Price: 25.00},
				{ID: 3, Name: "Keyboard", Price: 75.00},
			}
			c.JSON(http.StatusOK, products)
		})

		// Simulate slow API
		api.GET("/slow", func(c *gin.Context) {
			// time.Sleep(3 * time.Second)
			c.JSON(http.StatusOK, gin.H{"message": "This was a slow endpoint"})
		})

		// Simulate error
		api.GET("/error", func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		})
	}

	log.Println("🚀 HTTP Mock Server starting on :17002")
	if err := router.Run(":17002"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
