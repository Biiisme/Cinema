package handler

import (
	"cinema/model"
	"cinema/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SeatHandler struct {
	SeatRepo repository.SeatRepo
}

func NewSeatHandler(repo repository.SeatRepo) *SeatHandler {
	return &SeatHandler{SeatRepo: repo}
}

func (u *SeatHandler) GetSeatbyFilmID(c *gin.Context) {
	id := c.Param("id")
	roomID, err := strconv.Atoi(id)
	seats, err := u.SeatRepo.GetAllSeat(c.Request.Context(), roomID)
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
		Message:    "Lấy tất cả ghế theo phòng thành công",
		Data:       seats,
	})
}
