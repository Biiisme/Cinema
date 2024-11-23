package repository

import (
	"cinema/model"
	"context"
)

// ScheduleRepo là interface cho các hành động liên quan đến Schedule
type ScheduleRepo interface {
	SaveSchedule(ctx context.Context, schedule model.Schedule) (model.Schedule, error)

	GetSchedulesByMovieID(ctx context.Context, movieID int) ([]model.Schedule, error)
}
