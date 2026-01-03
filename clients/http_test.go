package clients

import (
	"encoding/json"
	"testing"
)

// TestHTTPClient_Get tests HTTP GET requests to mock server
func TestHTTPClient_Get(t *testing.T) {
	// このテストはdocker-composeでhttp-mockサーバーが起動している必要があります
	// docker-compose up -d http-mock
	client := NewHTTPClient("http://localhost:17002")

	body, err := client.Get("/api/v1/users")
	if err != nil {
		t.Skipf("Mock server not available: %v", err)
		return
	}

	var users []map[string]interface{}
	if err := json.Unmarshal(body, &users); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(users) == 0 {
		t.Error("Expected at least one user")
	}

	t.Logf("Got %d users from mock server", len(users))
}

// TestHTTPClient_Post tests HTTP POST requests to mock server
func TestHTTPClient_Post(t *testing.T) {
	client := NewHTTPClient("http://localhost:17002")

	userData := map[string]interface{}{
		"name":      "Test User",
		"email":     "test@example.com",
		"user_type": "test",
	}

	body, err := client.Post("/api/v1/users", userData)
	if err != nil {
		t.Skipf("Mock server not available: %v", err)
		return
	}

	var user map[string]interface{}
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if user["id"] == nil {
		t.Error("Expected user ID in response")
	}

	t.Logf("Created user: %+v", user)
}
