package repository

import (
	"cinema/model"
	"context"
)

type CinemaRepo interface {
	GetAllCinemas(ctx context.Context) ([]model.Cinemas, error)
}
