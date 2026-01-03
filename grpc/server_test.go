package grpc

import (
	"context"
	"golang-gin/models"
	"testing"

	pb "golang-gin/grpc/proto"
)

func TestServer_GetAlbums(t *testing.T) {
	server := NewServer()
	ctx := context.Background()

	resp, err := server.GetAlbums(ctx, &pb.GetAlbumsRequest{})
	if err != nil {
		t.Fatalf("GetAlbums failed: %v", err)
	}

	if len(resp.Albums) == 0 {
		t.Error("Expected at least one album")
	}

	// Verify first album matches models.Albums
	if len(resp.Albums) > 0 && len(models.Albums) > 0 {
		if resp.Albums[0].Id != models.Albums[0].ID {
			t.Errorf("Expected album ID %s, got %s", models.Albums[0].ID, resp.Albums[0].Id)
		}
	}

	t.Logf("GetAlbums returned %d albums", len(resp.Albums))
}

func TestServer_GetAlbumByID(t *testing.T) {
	server := NewServer()
	ctx := context.Background()

	tests := []struct {
		name      string
		id        string
		expectNil bool
	}{
		{"Valid ID", "1", false},
		{"Invalid ID", "999", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := server.GetAlbumByID(ctx, &pb.GetAlbumByIDRequest{Id: tt.id})
			if err != nil {
				t.Fatalf("GetAlbumByID failed: %v", err)
			}

			if tt.expectNil && resp != nil {
				t.Error("Expected nil response for invalid ID")
			}
			if !tt.expectNil && resp == nil {
				t.Error("Expected non-nil response for valid ID")
			}

			if resp != nil && resp.Id != tt.id {
				t.Errorf("Expected album ID %s, got %s", tt.id, resp.Id)
			}
		})
	}
}

func TestServer_CreateAlbum(t *testing.T) {
	server := NewServer()
	ctx := context.Background()

	initialCount := len(models.Albums)

	req := &pb.CreateAlbumRequest{
		Id:     "999",
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  29.99,
		Tax:    0.1,
	}

	resp, err := server.CreateAlbum(ctx, req)
	if err != nil {
		t.Fatalf("CreateAlbum failed: %v", err)
	}

	if resp.Id != req.Id {
		t.Errorf("Expected album ID %s, got %s", req.Id, resp.Id)
	}
	if resp.Title != req.Title {
		t.Errorf("Expected title %s, got %s", req.Title, resp.Title)
	}
	if resp.Artist != req.Artist {
		t.Errorf("Expected artist %s, got %s", req.Artist, resp.Artist)
	}

	// Verify album was added to models.Albums
	if len(models.Albums) != initialCount+1 {
		t.Errorf("Expected %d albums, got %d", initialCount+1, len(models.Albums))
	}

	t.Logf("Created album: %+v", resp)
}
