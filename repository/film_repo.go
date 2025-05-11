package repository

import (
	"cinema/model"
	"cinema/model/req"
	"context"
)

type FilmRepo interface {
	SaveFilm(ctx context.Context, film model.Film) (model.Film, error)
	GetFilmByID(ctx context.Context, id string) (model.Film, error)
	GetAllFilms(ctx context.Context, offset int, length int) ([]model.Film, error)
	Delete(film model.Film) error
	UpdateFilm(filmReq req.FilmReq, id string) (model.Film, error)
	TotalPage(film model.Film, length int) int
}
