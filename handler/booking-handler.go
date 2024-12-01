package handler

import (
	"cinema/model"
	"cinema/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	BookingRepo repository.BookingRepo
}

func NewBookingHandler(bookingRepo repository.BookingRepo) *BookingHandler {
	return &BookingHandler{BookingRepo: bookingRepo}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req struct {
		ScheduleID uint   `json:"scheduleId" binding:"required"`
		SeatIDs    []uint `json:"seatIds" binding:"required"`
		UserID     uint   `json:"userId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lưu từng ghế đã chọn
	for _, seatID := range req.SeatIDs {
		booking := model.Booking{
			UserID:     req.UserID,
			ScheduleID: req.ScheduleID,
			SeatID:     seatID,
			Status:     "booked",
		}
		if err := h.BookingRepo.SaveBooking(booking); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book seat"})
			return
		}
	}

	// Lấy hóa đơn
	invoice, err := h.BookingRepo.GetBookingInvoice(req.ScheduleID, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve booking invoice"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking successful", "invoice": invoice})
}
