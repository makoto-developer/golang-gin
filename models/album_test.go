package models

import (
	"encoding/json"
	"testing"
)

func TestAlbum_Struct(t *testing.T) {
	album := Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  19.99,
		Tax:    0.1,
	}

	if album.ID != "1" {
		t.Errorf("Expected ID '1', got '%s'", album.ID)
	}
	if album.Title != "Test Album" {
		t.Errorf("Expected Title 'Test Album', got '%s'", album.Title)
	}
	if album.Artist != "Test Artist" {
		t.Errorf("Expected Artist 'Test Artist', got '%s'", album.Artist)
	}
	if album.Price != 19.99 {
		t.Errorf("Expected Price 19.99, got %f", album.Price)
	}
	if album.Tax != 0.1 {
		t.Errorf("Expected Tax 0.1, got %f", album.Tax)
	}
}

func TestAlbum_JSONMarshaling(t *testing.T) {
	album := Album{
		ID:     "1",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  19.99,
		Tax:    0.1,
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(album)
	if err != nil {
		t.Fatalf("Failed to marshal album: %v", err)
	}

	// Unmarshal back to struct
	var unmarshaled Album
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal album: %v", err)
	}

	// Compare
	if unmarshaled.ID != album.ID {
		t.Errorf("ID mismatch: expected '%s', got '%s'", album.ID, unmarshaled.ID)
	}
	if unmarshaled.Title != album.Title {
		t.Errorf("Title mismatch: expected '%s', got '%s'", album.Title, unmarshaled.Title)
	}
	if unmarshaled.Artist != album.Artist {
		t.Errorf("Artist mismatch: expected '%s', got '%s'", album.Artist, unmarshaled.Artist)
	}
	if unmarshaled.Price != album.Price {
		t.Errorf("Price mismatch: expected %f, got %f", album.Price, unmarshaled.Price)
	}
	if unmarshaled.Tax != album.Tax {
		t.Errorf("Tax mismatch: expected %f, got %f", album.Tax, unmarshaled.Tax)
	}
}

func TestAlbums_InitialData(t *testing.T) {
	if len(Albums) == 0 {
		t.Error("Expected Albums to have initial data")
	}

	// Verify first album
	if len(Albums) > 0 {
		firstAlbum := Albums[0]
		if firstAlbum.ID == "" {
			t.Error("First album should have an ID")
		}
		if firstAlbum.Title == "" {
			t.Error("First album should have a title")
		}
		if firstAlbum.Artist == "" {
			t.Error("First album should have an artist")
		}
	}

	t.Logf("Albums has %d initial records", len(Albums))
}

func TestAlbum_JSONTags(t *testing.T) {
	jsonStr := `{
		"id": "test-id",
		"title": "Test Title",
		"artist": "Test Artist",
		"price": 29.99,
		"tax": 0.08
	}`

	var album Album
	err := json.Unmarshal([]byte(jsonStr), &album)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if album.ID != "test-id" {
		t.Errorf("Expected ID 'test-id', got '%s'", album.ID)
	}
	if album.Title != "Test Title" {
		t.Errorf("Expected Title 'Test Title', got '%s'", album.Title)
	}
	if album.Artist != "Test Artist" {
		t.Errorf("Expected Artist 'Test Artist', got '%s'", album.Artist)
	}
	if album.Price != 29.99 {
		t.Errorf("Expected Price 29.99, got %f", album.Price)
	}
	if album.Tax != 0.08 {
		t.Errorf("Expected Tax 0.08, got %f", album.Tax)
	}
}
