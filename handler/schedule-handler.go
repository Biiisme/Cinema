package handler

import (
	"cinema/model"
	"cinema/model/req"
	"cinema/repository/repo_impl"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct {
	ScheduleRepo repo_impl.ScheduleRepoImpl
}

func (h *ScheduleHandler) HandleSaveSchedule(c *gin.Context) {
	var scheduleReq req.ReqSchedule
	if err := c.ShouldBindJSON(&scheduleReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	layout := "2006-01-02" // Định dạng ngày
	ShowDate, err := time.Parse(layout, scheduleReq.ShowDate)
	if err != nil {
		// Xử lý lỗi nếu không thể chuyển đổi
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	layoutTime := "15:04:05" // Định dạng thời gian
	ShowTime, err := time.Parse(layoutTime, scheduleReq.ShowTime)
	if err != nil {
		// Xử lý lỗi nếu không thể chuyển đổi
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}
	schedule := model.Schedule{
		FilmID:   scheduleReq.FilmID,
		RoomID:   scheduleReq.RoomID,
		ShowDate: ShowDate,
		ShowTime: ShowTime,
	}

	schedule, err = h.ScheduleRepo.SaveSchedule(c.Request.Context(), schedule)
	if err != nil {
		c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Lưu ngày chiếu thành công",
		Data:       schedule,
	})
}

// HandleGetSchedulesByFilmID xử lý lấy danh sách lịch chiếu theo FilmID
func (h *ScheduleHandler) HandleGetSchedulesByFilmID(c *gin.Context) {
	id := c.Param("filmId") // Lấy FilmID từ URL param

	// Chuyển ID từ string sang int
	filmID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID"})
		return
	}

	// Lấy danh sách lịch chiếu từ repository
	schedules, err := h.ScheduleRepo.GetSchedulesByMovieID(c.Request.Context(), filmID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving schedules", "details": err.Error()})
		return
	}

	// Trả về danh sách lịch chiếu
	c.JSON(http.StatusOK, gin.H{
		"message": "Schedules retrieved successfully",
		"data":    schedules,
	})
}