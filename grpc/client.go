package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "golang-gin/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client wraps the gRPC client connection
type Client struct {
	conn   *grpc.ClientConn
	client pb.AlbumServiceClient
}

// NewClient creates a new gRPC client
func NewClient(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	client := pb.NewAlbumServiceClient(conn)
	return &Client{
		conn:   conn,
		client: client,
	}, nil
}

// Close closes the gRPC connection
func (c *Client) Close() error {
	return c.conn.Close()
}

// GetAlbums retrieves all albums via gRPC
func (c *Client) GetAlbums() (*pb.GetAlbumsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.GetAlbums(ctx, &pb.GetAlbumsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get albums: %w", err)
	}

	return resp, nil
}

// GetAlbumByID retrieves a specific album by ID via gRPC
func (c *Client) GetAlbumByID(id string) (*pb.Album, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.GetAlbumByID(ctx, &pb.GetAlbumByIDRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("failed to get album: %w", err)
	}

	return resp, nil
}

// CreateAlbum creates a new album via gRPC
func (c *Client) CreateAlbum(id, title, artist string, price float64, tax float32) (*pb.Album, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.CreateAlbum(ctx, &pb.CreateAlbumRequest{
		Id:     id,
		Title:  title,
		Artist: artist,
		Price:  price,
		Tax:    tax,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create album: %w", err)
	}

	return resp, nil
}

// Example demonstrates how to use the gRPC client
func Example() {
	client, err := NewClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Get all albums
	albums, err := client.GetAlbums()
	if err != nil {
		log.Fatalf("Failed to get albums: %v", err)
	}
	log.Printf("Albums: %v", albums)

	// Get album by ID
	album, err := client.GetAlbumByID("1")
	if err != nil {
		log.Fatalf("Failed to get album: %v", err)
	}
	log.Printf("Album: %v", album)

	// Create album
	newAlbum, err := client.CreateAlbum("4", "New Album", "Artist", 29.99, 0.1)
	if err != nil {
		log.Fatalf("Failed to create album: %v", err)
	}
	log.Printf("Created album: %v", newAlbum)
}
