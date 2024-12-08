package repo_impl

import (
	"cinema/banana"
	"cinema/model"
	"cinema/repository"
	"context"
	"errors"
	"log"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type FilmRepoImpl struct {
	db *gorm.DB
}

// NewFilmRepoImpl create NewFilmRepoImpl
func NewFilmRepoImpl(db *gorm.DB) repository.FilmRepo {
	return &FilmRepoImpl{
		db: db,
	}
}

// SaveFilm to database
func (f *FilmRepoImpl) SaveFilm(ctx context.Context, film model.Film) (model.Film, error) {
	// Using GORM to same film
	if err := f.db.WithContext(ctx).Create(&film).Error; err != nil {
		// Check error is unique constraint for PostgreSQL
		var pqErr *pq.Error
		if ok := errors.As(err, &pqErr); ok && pqErr.Code.Name() == "unique_violation" {
			return film, banana.FilmConflict
		}
		log.Println("Error saving film:", err)
		return film, banana.SaveFilmFail
	}

	return film, nil
}

// GetFilmByID
func (f *FilmRepoImpl) GetFilmByID(ctx context.Context, id string) (model.Film, error) {
	var film model.Film
	// Check film by ID (GORM)
	if err := f.db.WithContext(ctx).First(&film, "film_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return film, banana.FilmNotFound
		}
		log.Println("Error retrieving film by ID:", err)
		return film, err
	}

	return film, nil
}

// GetAllFilms for database
func (f *FilmRepoImpl) GetAllFilms(ctx context.Context) ([]model.Film, error) {
	var films []model.Film
	// Get all films
	if err := f.db.WithContext(ctx).Find(&films).Error; err != nil {
		log.Println("Error retrieving all films:", err)
		return nil, err
	}

	return films, nil
}
