package handler

import (
	"cinema/model"
	"cinema/model/req"
	"cinema/repository"
	"cinema/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	TicketRepo repository.TicketRepo
	SeatRepo   repository.SeatRepo
}

func NewBookingHandler(ticketRepo repository.TicketRepo, seatRepo repository.SeatRepo) *TicketHandler {
	return &TicketHandler{TicketRepo: ticketRepo, SeatRepo: seatRepo}

}

func (h *TicketHandler) CreateTicket(c *gin.Context) {
	var bookingReq req.BookingTicketRequest
	if err := c.ShouldBindJSON(&bookingReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	key := "booking:session:" + bookingReq.BookingSessionID

	val, _ := utils.Cache.Get(key).Result()

	var data req.HoldSeatRequest
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid session data"})
		return
	}

	userID, _ := c.Get("user_id")
	userIDUstring, _ := userID.(string)

	totalprice := len(data.Seats) * 45000
	seat := strings.Join(data.SeatName, ",")
	ticket := model.Ticket{
		UserID:      userIDUstring,
		ScheduleID:  int32(data.ScheduleId),
		Seat:        seat,
		Code:        123,
		TotalPrice:  int64(totalprice),
		Status:      "booked",
		PaymentDate: time.Now(),
	}
	err := h.TicketRepo.SaveTicket(ticket)
	if err != nil {
		c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	if h.SeatRepo == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Seat repository not initialized",
		})
		return
	}
	for _, seatId := range data.Seats {
		err := h.SeatRepo.UpdateStatusSeat(c, seatId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Mua vé không thành công",
				Data:       nil,
			})
			return
		}
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Mua vé thành công",
		Data:       ticket,
	})
}

func (h *TicketHandler) HoldSeat(c *gin.Context) {
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
	now := time.Now()
	req.CreatedAt = &now

	jsonData, _ := json.Marshal(req)

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

func (h *TicketHandler) GetHoldSeatInfo(c *gin.Context) {
	sessionId := c.Query("bookingSessionId")

	key := "booking:session:" + sessionId

	val, _ := utils.Cache.Get(key).Result()

	var data req.HoldSeatRequest
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid session data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h *TicketHandler) FindTicket(c *gin.Context) {

	userID, _ := c.Get("user_id")
	userIDUstring, _ := userID.(string)

	tickets, err := h.TicketRepo.FindTicket(userIDUstring)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Lấy danh sách vé mua thành công thành công",
		Data:       tickets,
	})
}
