package repo_impl

import (
	"cinema/model"
	"cinema/repository"
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type SeatRepoImpl struct {
	db *gorm.DB
}

func NewSeatRepoImpl(db *gorm.DB) repository.SeatRepo {
	return &SeatRepoImpl{
		db: db,
	}
}

// GetAllCinemas implements repository.CinemaRepo.
func (c *SeatRepoImpl) GetAllSeat(ctx context.Context, room_id int) ([]model.Seat, error) {
	var seats []model.Seat
	// Get all films
	if err := c.db.WithContext(ctx).Where("room_id=?", room_id).Find(&seats).Error; err != nil {
		log.Println("Error retrieving all cinema:", err)
		return nil, err
	}

	return seats, nil
}

func (c *SeatRepoImpl) UpdateStatusSeat(ctx context.Context, SeatId int) error {
	fmt.Println("run1")
	result := c.db.Model(&model.Seat{}).Where("id = ?", SeatId).Updates(map[string]interface{}{"status": "booked"})

	return result.Error
}
