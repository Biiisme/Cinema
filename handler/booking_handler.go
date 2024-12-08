package handler

import (
	"cinema/model"
	"cinema/model/req"
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
	var req req.ReqBooking

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	userIDUstring, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	// Save booking
	for _, seatID := range req.SeatID {
		booking := model.Booking{
			UserID:     userIDUstring,
			ScheduleID: req.ScheduleID,
			SeatID:     seatID,
			Status:     "booked",
		}
		if err := h.BookingRepo.SaveBooking(booking); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book seat"})
			return
		}
	}

	// Get Bill
	invoice, err := h.BookingRepo.GetBookingInvoice(req.ScheduleID, userIDUstring)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve booking invoice"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking successful", "invoice": invoice})
}
