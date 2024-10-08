package repo_impl

import (
	"cinema/banana"
	"cinema/db"
	"cinema/model"
	"cinema/repository"
	"context"
	"database/sql"

	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
)

type FilmRepoImpl struct {
	sql *db.Sql
}

// NewFilmRepoImpl tạo một đối tượng FilmRepoImpl mới
func NewFilmRepoImpl(sql *db.Sql) repository.FilmRepo {
	return &FilmRepoImpl{
		sql: sql,
	}
}

// SaveFilm lưu phim vào cơ sở dữ liệu
func (f FilmRepoImpl) SaveFilm(ctx context.Context, film model.Film) (model.Film, error) {
	statement := `
        INSERT INTO films(film_id, film_name, timefull, limitAge, image)
        VALUES(:film_id, :film_name, :timefull, :limitAge, :image)
    `

	_, err := f.sql.Db.NamedExecContext(ctx, statement, film)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return film, banana.FilmConflict
			}
		}
		return film, banana.SaveFilmFail
	}

	return film, nil
}

// GetFilmByID lấy thông tin phim theo ID
func (f FilmRepoImpl) GetFilmByID(ctx context.Context, id string) (model.Film, error) {
	var film model.Film
	err := f.sql.Db.GetContext(ctx, &film, "SELECT * FROM films WHERE film_id=$1", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return film, banana.FilmNotFound
		}
		log.Error(err.Error())
		return film, err
	}

	return film, nil
}

// GetAllFilms lấy tất cả phim từ cơ sở dữ liệu
func (f FilmRepoImpl) GetAllFilms(ctx context.Context) ([]model.Film, error) {
	var films []model.Film
	err := f.sql.Db.SelectContext(ctx, &films, "SELECT * FROM films")

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return films, nil
}
