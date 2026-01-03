package grpc

import (
	"context"
	"golang-gin/repository"
	"testing"

	pb "golang-gin/grpc/proto"
)

func TestServer_GetAlbums(t *testing.T) {
	mockRepo := repository.NewMockAlbumRepository()
	server := NewServer(mockRepo)
	ctx := context.Background()

	resp, err := server.GetAlbums(ctx, &pb.GetAlbumsRequest{})
	if err != nil {
		t.Fatalf("GetAlbums failed: %v", err)
	}

	if len(resp.Albums) == 0 {
		t.Error("Expected at least one album")
	}

	if len(resp.Albums) > 0 {
		if resp.Albums[0].Id != "1" {
			t.Errorf("Expected album ID 1, got %s", resp.Albums[0].Id)
		}
	}

	t.Logf("GetAlbums returned %d albums", len(resp.Albums))
}

func TestServer_GetAlbumByID(t *testing.T) {
	mockRepo := repository.NewMockAlbumRepository()
	server := NewServer(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name      string
		id        string
		expectNil bool
	}{
		{"Valid ID", "1", false},
		{"Invalid ID", "999", true},
		{"Invalid ID format", "abc", false}, // エラーが返されるが nil にはならない
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := server.GetAlbumByID(ctx, &pb.GetAlbumByIDRequest{Id: tt.id})

			if tt.name == "Invalid ID format" {
				if err == nil {
					t.Error("Expected error for invalid ID format")
				}
				return
			}

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
	mockRepo := repository.NewMockAlbumRepository()
	server := NewServer(mockRepo)
	ctx := context.Background()

	req := &pb.CreateAlbumRequest{
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  29.99,
		Tax:    0.1,
	}

	resp, err := server.CreateAlbum(ctx, req)
	if err != nil {
		t.Fatalf("CreateAlbum failed: %v", err)
	}

	if resp.Id == "" {
		t.Error("Expected non-empty album ID")
	}
	if resp.Title != req.Title {
		t.Errorf("Expected title %s, got %s", req.Title, resp.Title)
	}
	if resp.Artist != req.Artist {
		t.Errorf("Expected artist %s, got %s", req.Artist, resp.Artist)
	}

	t.Logf("Created album: %+v", resp)
}
