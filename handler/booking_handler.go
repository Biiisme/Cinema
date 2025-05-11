package handler

import (
	"cinema/model/req"
	"cinema/repository"
	"cinema/utils"
	"encoding/json"
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

// func (h *BookingHandler) CreateBooking(c *gin.Context) {
// 	var req req.HoldSeatRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	userID, _ := c.Get("user_id")
// 	userIDUstring, _ := userID.(string)

// 	// Save booking
// 	for _, seatID := range req.SeatID {
// 		booking := model.Booking{
// 			UserID:     userIDUstring,
// 			FilmID:     req.FilmID,
// 			ScheduleID: req.ScheduleID,
// 			SeatID:     seatID,
// 			Status:     "booked",
// 		}
// 		if err := h.BookingRepo.SaveBooking(booking); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book seat"})
// 			return
// 		}
// 	}

// 	// Get Bill
// 	invoice, err := h.BookingRepo.GetBookingInvoice(req.ScheduleID, userIDUstring)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve booking invoice"})
// 		return
// 	}

//		c.JSON(http.StatusOK, gin.H{"message": "Booking successful", "invoice": invoice})
//	}
func (h *BookingHandler) HoldSeat(c *gin.Context) {
	var req req.HoldSeatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sessionId, _ := c.Get("user_id")

	for _, seatId := range req.Seats {
		strSeat := fmt.Sprintf("seat_id:%s", seatId)
		ok, err := utils.Cache.SetNX(strSeat, sessionId, 5*time.Minute).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi Redis"})
			return
		}
		if !ok {
			c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("Ghế %d đang được giữ bởi người khác", seatId)})
			return
		}
	}

	data := map[string]interface{}{
		"film_id":     req.FilmId,
		"schedule_id": req.ScheduleId,
		"seats":       req.Seats,
		"createdAt":   time.Now().Format(time.RFC3339),
	}

	jsonData, _ := json.Marshal(data)

	key := "booking:session:" + sessionId.(string)
	err := utils.Cache.Set(key, jsonData, 5*time.Minute).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to hold seat"})
		return
	}

	c.JSON(200, gin.H{
		"bookingSessionId": sessionId,
	})
}

func (h *BookingHandler) GetHoldSeatInfo(c *gin.Context) {
	sessionId := c.Query("bookingSessionId")

	key := "booking:session:" + sessionId

	val, _ := utils.Cache.Get(key).Result()

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid session data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
