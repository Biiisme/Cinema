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

// NewFilmRepoImpl tạo một đối tượng FilmRepoImpl mới
func NewFilmRepoImpl(db *gorm.DB) repository.FilmRepo {
	return &FilmRepoImpl{
		db: db,
	}
}

// SaveFilm lưu phim vào cơ sở dữ liệu
func (f *FilmRepoImpl) SaveFilm(ctx context.Context, film model.Film) (model.Film, error) {
	// Sử dụng GORM để lưu phim
	if err := f.db.WithContext(ctx).Create(&film).Error; err != nil {
		// Kiểm tra xem lỗi có phải là unique constraint từ PostgreSQL
		var pqErr *pq.Error
		if ok := errors.As(err, &pqErr); ok && pqErr.Code.Name() == "unique_violation" {
			return film, banana.FilmConflict
		}
		log.Println("Error saving film:", err)
		return film, banana.SaveFilmFail
	}

	return film, nil
}

// GetFilmByID lấy thông tin phim theo ID
func (f *FilmRepoImpl) GetFilmByID(ctx context.Context, id string) (model.Film, error) {
	var film model.Film
	// Tìm phim theo ID bằng GORM
	if err := f.db.WithContext(ctx).First(&film, "film_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return film, banana.FilmNotFound
		}
		log.Println("Error retrieving film by ID:", err)
		return film, err
	}

	return film, nil
}

// GetAllFilms lấy tất cả phim từ cơ sở dữ liệu
func (f *FilmRepoImpl) GetAllFilms(ctx context.Context) ([]model.Film, error) {
	var films []model.Film
	// Lấy tất cả các bản ghi từ bảng films
	if err := f.db.WithContext(ctx).Find(&films).Error; err != nil {
		log.Println("Error retrieving all films:", err)
		return nil, err
	}

	return films, nil
}
