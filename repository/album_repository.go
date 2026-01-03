package repository

import (
	"golang-gin/models"

	"gorm.io/gorm"
)

// AlbumRepository defines the interface for album data access
type AlbumRepository interface {
	FindAll() ([]models.Album, error)
	FindByID(id uint) (*models.Album, error)
	Create(album *models.Album) error
	Update(album *models.Album) error
	Delete(id uint) error
}

// albumRepository implements AlbumRepository
type albumRepository struct {
	db *gorm.DB
}

// NewAlbumRepository creates a new AlbumRepository instance
func NewAlbumRepository(db *gorm.DB) AlbumRepository {
	return &albumRepository{db: db}
}

// FindAll retrieves all albums
func (r *albumRepository) FindAll() ([]models.Album, error) {
	var albums []models.Album
	if err := r.db.Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

// FindByID retrieves an album by ID
func (r *albumRepository) FindByID(id uint) (*models.Album, error) {
	var album models.Album
	if err := r.db.First(&album, id).Error; err != nil {
		return nil, err
	}
	return &album, nil
}

// Create creates a new album
func (r *albumRepository) Create(album *models.Album) error {
	return r.db.Create(album).Error
}

// Update updates an existing album
func (r *albumRepository) Update(album *models.Album) error {
	return r.db.Save(album).Error
}

// Delete soft deletes an album by ID
func (r *albumRepository) Delete(id uint) error {
	return r.db.Delete(&models.Album{}, id).Error
}
