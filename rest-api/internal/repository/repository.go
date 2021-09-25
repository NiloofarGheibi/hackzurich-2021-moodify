package repository

import (
	"context"
	"rest-api/internal/models"
)

type Repository interface {
	// GetByID fetches based on id
	GetByID(ctx context.Context, id string) (*models.User, error)
	// Get gets all in paginated fashion
	Get(ctx context.Context, filters *Filters, offset int, limit int) ([]*models.User, error)
	// StoreUser stores the input
	StoreUser(ctx context.Context, data *models.User) (*models.User, error)
	// StoreTrack stores the input
	StoreTrack(ctx context.Context, data *models.Track) (*models.Track, error)
	// Update updates existing data
	Update(ctx context.Context, data *models.User) (*models.User, error)
	// Delete deletes existing data
	Delete(ctx context.Context, data *models.User) error
}
