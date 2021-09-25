package services

import (
	"context"
	"rest-api/internal/models"
	"rest-api/internal/repository"
)

type User interface {
	// GetByID fetches based on id
	GetByID(ctx context.Context, id string) (*models.User, error)
	// GetByEmail fetches based on email
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	// Get gets all in paginated fashion
	Get(ctx context.Context, offset int, limit int) ([]*models.User, error)
	// Store stores the input
	Store(ctx context.Context, user *models.User) (*models.User, error)
	// Update updates existing data
	Update(ctx context.Context, user *models.User, updates *models.User) (*models.User, error)
	// Delete deletes existing data
	Delete(ctx context.Context, user *models.User) error
}

type user struct {
	Repo repository.Repository
}

func NewUserService(repo repository.Repository) *user {
	return &user{Repo: repo}
}

func (u user) GetByID(ctx context.Context, id string) (*models.User, error) {
	user, err := u.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u user) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.Repo.Get(ctx, &repository.Filters{Email: email}, 0, 1)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user[0], nil
}

func (u user) Get(ctx context.Context,offset int, limit int) ([]*models.User, error) {
	users, err := u.Repo.Get(ctx, &repository.Filters{}, offset, limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u user) Store(ctx context.Context, user *models.User) (*models.User, error) {
	newUser, err := u.Repo.StoreUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u user) Update(ctx context.Context, user *models.User, updates *models.User) (*models.User, error) {
	user.FirstName = updates.FirstName
	user.LastName = updates.LastName
	updatedUser, err := u.Repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u user) Delete(ctx context.Context, user *models.User) error {
	return u.Repo.Delete(ctx, user)
}
