package repo_impl

import (
	"cinema/model"
	"cinema/repository"

	"gorm.io/gorm"
)

type BookingRepoImpl struct {
	db *gorm.DB
}

func NewBookingRepo(db *gorm.DB) repository.BookingRepo {
	return &BookingRepoImpl{db: db}
}

func (r *BookingRepoImpl) SaveBooking(booking model.Booking) error {
	return r.db.Create(&booking).Error
}

func (r *BookingRepoImpl) GetBookingInvoice(scheduleID uint, userID uint) (model.BookingInvoice, error) {
	var invoice model.BookingInvoice

	err := r.db.Raw(`
		SELECT  u.name AS user_name, f.name AS film_name, 
			s.show_date AS show_date, 
			s.show_time AS show_time, 
			r.name AS room_name, 
			ARRAY_AGG(seat.seat_number) AS seat_numbers,
			SUM(seat.price) AS total_price
		FROM bookings b
		JOIN users u ON b.user_id = u.id
		JOIN schedules s ON b.schedule_id = s.id
		JOIN films f ON s.film_id = f.id
		JOIN rooms r ON s.room_id = r.id
		JOIN seats seat ON b.seat_id = seat.id
		WHERE b.schedule_id = ? AND b.user_id = ?
		GROUP BY u.name, f.name, s.show_date, s.show_time, r.name
	`, scheduleID, userID).Scan(&invoice).Error

	return invoice, err
}
