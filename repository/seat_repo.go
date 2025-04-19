package repository

import (
	"cinema/model"
	"context"
)

type SeatRepo interface {
	GetAllSeat(ctx context.Context, room_id int) ([]model.Seat, error)
}
