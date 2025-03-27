package repository

import (
	"cinema/model"
	"context"
)

type FilmRepo interface {
	SaveFilm(ctx context.Context, film model.Film) (model.Film, error)
	GetFilmByID(ctx context.Context, id string) (model.Film, error)
	GetAllFilms(ctx context.Context) ([]model.Film, error)
	Delete(film model.Film) error
}
