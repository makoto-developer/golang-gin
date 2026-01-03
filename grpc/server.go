package grpc

import (
	"context"
	"fmt"
	"golang-gin/models"
	"golang-gin/repository"
	"strconv"

	pb "golang-gin/grpc/proto"
	"gorm.io/gorm"
)

// Server implements the AlbumService gRPC server
type Server struct {
	pb.UnimplementedAlbumServiceServer
	repo repository.AlbumRepository
}

// NewServer creates a new gRPC server instance
func NewServer(repo repository.AlbumRepository) *Server {
	return &Server{repo: repo}
}

// GetAlbums returns all albums
func (s *Server) GetAlbums(ctx context.Context, req *pb.GetAlbumsRequest) (*pb.GetAlbumsResponse, error) {
	albums, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var pbAlbums []*pb.Album
	for _, a := range albums {
		pbAlbums = append(pbAlbums, &pb.Album{
			Id:     fmt.Sprintf("%d", a.ID),
			Title:  a.Title,
			Artist: a.Artist,
			Price:  a.Price,
			Tax:    a.Tax,
		})
	}
	return &pb.GetAlbumsResponse{Albums: pbAlbums}, nil
}

// GetAlbumByID returns a specific album by ID
func (s *Server) GetAlbumByID(ctx context.Context, req *pb.GetAlbumByIDRequest) (*pb.Album, error) {
	id, err := strconv.ParseUint(req.Id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid album ID: %w", err)
	}

	album, err := s.repo.FindByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &pb.Album{
		Id:     fmt.Sprintf("%d", album.ID),
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
		Tax:    album.Tax,
	}, nil
}

// CreateAlbum creates a new album
func (s *Server) CreateAlbum(ctx context.Context, req *pb.CreateAlbumRequest) (*pb.Album, error) {
	albumModel := &models.Album{
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
		Tax:    req.Tax,
	}

	if err := s.repo.Create(albumModel); err != nil {
		return nil, err
	}

	return &pb.Album{
		Id:     fmt.Sprintf("%d", albumModel.ID),
		Title:  albumModel.Title,
		Artist: albumModel.Artist,
		Price:  albumModel.Price,
		Tax:    albumModel.Tax,
	}, nil
}
