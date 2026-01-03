package grpc

import (
	"context"
	"golang-gin/models"

	pb "golang-gin/grpc/proto"
)

// Server implements the AlbumService gRPC server
type Server struct {
	pb.UnimplementedAlbumServiceServer
}

// NewServer creates a new gRPC server instance
func NewServer() *Server {
	return &Server{}
}

// GetAlbums returns all albums
func (s *Server) GetAlbums(ctx context.Context, req *pb.GetAlbumsRequest) (*pb.GetAlbumsResponse, error) {
	var albums []*pb.Album
	for _, a := range models.Albums {
		albums = append(albums, &pb.Album{
			Id:     a.ID,
			Title:  a.Title,
			Artist: a.Artist,
			Price:  a.Price,
			Tax:    a.Tax,
		})
	}
	return &pb.GetAlbumsResponse{Albums: albums}, nil
}

// GetAlbumByID returns a specific album by ID
func (s *Server) GetAlbumByID(ctx context.Context, req *pb.GetAlbumByIDRequest) (*pb.Album, error) {
	for _, a := range models.Albums {
		if a.ID == req.Id {
			return &pb.Album{
				Id:     a.ID,
				Title:  a.Title,
				Artist: a.Artist,
				Price:  a.Price,
				Tax:    a.Tax,
			}, nil
		}
	}
	return nil, nil
}

// CreateAlbum creates a new album
func (s *Server) CreateAlbum(ctx context.Context, req *pb.CreateAlbumRequest) (*pb.Album, error) {
	newAlbum := models.Album{
		ID:     req.Id,
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
		Tax:    req.Tax,
	}
	models.Albums = append(models.Albums, newAlbum)

	return &pb.Album{
		Id:     newAlbum.ID,
		Title:  newAlbum.Title,
		Artist: newAlbum.Artist,
		Price:  newAlbum.Price,
		Tax:    newAlbum.Tax,
	}, nil
}
