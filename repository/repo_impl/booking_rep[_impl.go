package repo_impl

import (
	"cinema/banana"
	"cinema/model"
	"context"
	"log"

	"gorm.io/gorm"
)

type BookingRepoImpl struct {
	db *gorm.DB
}

func NewBookingRepoImpl(db *gorm.DB) *BookingRepoImpl {
	return &BookingRepoImpl{db: db}
}

func (repo *BookingRepoImpl) SaveBooking(ctx context.Context, booking model.Booking) (model.Booking, error) {
	if err := repo.db.WithContext(ctx).Create(&booking).Error; err != nil {
		log.Println("Error saving schedule:", err)
		return booking, banana.SaveBooking
	}
	return booking, nil
}

func (repo *BookingRepoImpl) GetBookingsByScheduleID(scheduleID int) ([]model.Booking, error) {
	var bookings []model.Booking
	if err := repo.db.Where("schedule_id = ?", scheduleID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
