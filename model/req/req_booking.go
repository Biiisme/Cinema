package req

type HoldSeatRequest struct {
	FilmId     int   `json:"film_id" binding:"required"`
	ScheduleId int   `json:"schedule_id" binding:"required"`
	Seats      []int `json:"seats" binding:"required"` // Danh sách ghế chi tiết
}
