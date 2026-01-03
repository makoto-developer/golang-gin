package repository

import (
	"golang-gin/models"

	"gorm.io/gorm"
)

// MockAlbumRepository is a mock implementation of AlbumRepository for testing
type MockAlbumRepository struct {
	albums []models.Album
}

// NewMockAlbumRepository creates a new MockAlbumRepository with sample data
func NewMockAlbumRepository() *MockAlbumRepository {
	return &MockAlbumRepository{
		albums: []models.Album{
			{ID: 1, Title: "Hammerhead", Artist: "THE OFFSPRING", Price: 25.05, Tax: 0.1},
			{ID: 2, Title: "Shake It Off", Artist: "Taylor Swift", Price: 23.14, Tax: 0.1},
			{ID: 3, Title: "mysterious love", Artist: "Miho Komatsu", Price: 18.88, Tax: 0.1},
		},
	}
}

// FindAll retrieves all albums
func (m *MockAlbumRepository) FindAll() ([]models.Album, error) {
	return m.albums, nil
}

// FindByID retrieves an album by ID
func (m *MockAlbumRepository) FindByID(id uint) (*models.Album, error) {
	for _, album := range m.albums {
		if album.ID == id {
			return &album, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

// Create creates a new album
func (m *MockAlbumRepository) Create(album *models.Album) error {
	album.ID = uint(len(m.albums) + 1)
	m.albums = append(m.albums, *album)
	return nil
}

// Update updates an existing album
func (m *MockAlbumRepository) Update(album *models.Album) error {
	for i, a := range m.albums {
		if a.ID == album.ID {
			m.albums[i] = *album
			return nil
		}
	}
	return gorm.ErrRecordNotFound
}

// Delete soft deletes an album by ID
func (m *MockAlbumRepository) Delete(id uint) error {
	for i, album := range m.albums {
		if album.ID == id {
			m.albums = append(m.albums[:i], m.albums[i+1:]...)
			return nil
		}
	}
	return gorm.ErrRecordNotFound
}
