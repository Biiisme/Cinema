package repository

import (
	"cinema/model"
	"context"
)

type SeatRepo interface {
	GetAllSeat(ctx context.Context, room_id int) ([]model.Seat, error)
	UpdateStatusSeat(ctx context.Context, seat_id int) error
}
