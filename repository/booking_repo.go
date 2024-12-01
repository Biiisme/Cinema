package repository

import (
	"cinema/model"
)

type BookingRepo interface {
	SaveBooking(booking model.Booking) error
	GetBookingInvoice(scheduleID uint, userID uint) (model.BookingInvoice, error)
}
