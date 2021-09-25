package repository

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/rs/zerolog/log"
	"rest-api/internal/models"
	"time"
)

// Filters are used for filtering users
type Filters struct {
	Email string
}

type database struct {
	DB *pg.DB
}

func InitDatabaseConnection(url, username, password string) *pg.DB {
	opt, err := pg.ParseURL(url)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("unable to parse database url")
	}

	opt.User = username
	opt.Password = password
	return pg.Connect(opt)
}

func NewDatabaseRepository(db *pg.DB) *database {
	return &database{DB: db}
}

func (d database) GetByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{ID: id}

	err := d.DB.WithContext(ctx).Model(user).WherePK().Select()
	if err != nil {
		switch err {
		case pg.ErrNoRows:
			{
				return nil, nil
			}
		default:
			return nil, err
		}
	}

	return user, nil
}

func (d database) Get(ctx context.Context, filters *Filters, offset int, limit int) ([]*models.User, error) {
	var users []*models.User
	q := d.DB.WithContext(ctx).Model(&users)

	if filters.Email != "" {
		q = q.Where("email = ?", filters.Email)
	}

	if limit > 0 {
		q = q.Limit(limit)
	}

	err := q.Offset(offset).Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (d database) StoreUser(ctx context.Context, data *models.User) (*models.User, error) {
	_, err := d.DB.WithContext(ctx).Model(data).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d database) StoreTrack(ctx context.Context, data *models.Track) (*models.Track, error) {
	data.CreatedAt = time.Now().UTC()
	_, err := d.DB.WithContext(ctx).Model(data).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	//TODO
	fmt.Println(data.TrackID)
	return data, nil
}

func (d database) Update(ctx context.Context, data *models.User) (*models.User, error) {
	_, err := d.DB.WithContext(ctx).Model(data).WherePK().Update()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d database) Delete(ctx context.Context, data *models.User) error {
	_, err := d.DB.WithContext(ctx).Model(data).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}
