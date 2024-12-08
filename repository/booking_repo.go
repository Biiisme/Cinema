package repository

import (
	"cinema/model"
)

type BookingRepo interface {
	SaveBooking(booking model.Booking) error
	GetBookingInvoice(scheduleID int, userID string) (model.BookingInvoice, error)
}
