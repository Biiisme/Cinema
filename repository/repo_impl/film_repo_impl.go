package repo_impl

import (
	"cinema/banana"
	"cinema/model"
	"cinema/model/req"
	"cinema/repository"
	"context"
	"errors"
	"log"
	"time"

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
	if err := f.db.WithContext(ctx).First(&film, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return film, banana.FilmNotFound
		}
		log.Println("Error retrieving film by ID:", err)
		return film, err
	}

	return film, nil
}

// GetAllFilms for database
func (f *FilmRepoImpl) GetAllFilms(ctx context.Context, offset int, length int) ([]model.Film, error) {
	var films []model.Film
	// Get all films
	if err := f.db.WithContext(ctx).Limit(length).Offset(offset).Find(&films).Error; err != nil {
		log.Println("Error retrieving all films:", err)
		return nil, err
	}

	return films, nil
}

// Delete Film for database
func (f *FilmRepoImpl) Delete(film model.Film) error {
	err := f.db.Unscoped().Delete(&film).Error
	if err != nil {
		log.Println("Error delete film")
		return err
	}
	return nil
}

func (f *FilmRepoImpl) UpdateFilm(filmReq req.FilmReq, id string) (model.Film, error) {
	// Using GORM to same film
	var film model.Film
	if err := f.db.First(&film, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return film, banana.FilmNotFound
		}
		log.Println("Error retrieving film by ID:", err)
		return film, err
	}

	// Cập nhật các trường từ FilmReq vào Film
	updateData := map[string]interface{}{
		"title":          filmReq.Title,
		"poster_url":     filmReq.PosterURL,
		"trailer_url":    filmReq.TrailerURL,
		"description":    filmReq.Description,
		"duration":       filmReq.Duration,
		"release_date":   filmReq.ReleaseDate,
		"end_date":       filmReq.EndDate,
		"genre":          filmReq.Genre,
		"director":       filmReq.Director,
		"actors":         filmReq.Actors,
		"rated":          filmReq.Rated,
		"is_now_showing": filmReq.IsNowShowing,
		"is_coming_soon": filmReq.IsComingSoon,
		"updated_at":     time.Now(),
	}
	if err := f.db.Model(&film).Updates(updateData).Error; err != nil {
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

func (f *FilmRepoImpl) TotalPage(film model.Film, length int) int {
	var total int64

	f.db.Model(&film).Count((&total))

	totalPage := int((total + int64(length) - 1) / int64(length))
	return totalPage
}
