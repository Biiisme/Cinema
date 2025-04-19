package handler

import (
	"cinema/model"
	"cinema/model/req"
	"cinema/repository"
	"cinema/utils"
	"fmt"
	"net/http"
	"time"

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

	userID, _ := c.Get("user_id")
	userIDUstring, _ := userID.(string)

	// Save booking
	for _, seatID := range req.SeatID {
		booking := model.Booking{
			UserID:     userIDUstring,
			FilmID:     req.FilmID,
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
func (h *BookingHandler) CheckSeat(c *gin.Context) {
	var req req.ReqBooking

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashToken := "12345677"
	// Set redis
	for _, seatID := range req.SeatID {
		strSeat := fmt.Sprintf("seatID:%s", seatID)
		ok, err := utils.Cache.SetNX(strSeat, hashToken, 5*time.Minute).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi Redis"})
			return
		}
		if !ok {
			c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("Ghế %d đang được giữ bởi người khác", seatID)})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Giữ ghế thành công"})
}
