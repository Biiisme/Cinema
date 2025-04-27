package repo_impl

import (
	"cinema/model"
	"cinema/repository"
	"context"
	"log"

	"gorm.io/gorm"
)

type CinemaRepoImpl struct {
	db *gorm.DB
}

func NewCinemaRepoImpl(db *gorm.DB) repository.CinemaRepo {
	return &CinemaRepoImpl{
		db: db,
	}
}

// GetAllCinemas implements repository.CinemaRepo.
func (c *CinemaRepoImpl) GetAllCinemas(ctx context.Context) ([]model.Cinemas, error) {
	var cinemas []model.Cinemas
	// Get all films
	if err := c.db.WithContext(ctx).Find(&cinemas).Error; err != nil {
		log.Println("Error retrieving all cinema:", err)
		return nil, err
	}

	return cinemas, nil
}
