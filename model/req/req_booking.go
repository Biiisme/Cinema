package req

type ReqBooking struct {
	ScheduleID int   `json:"schedule_id" binding:"required"`
	SeatID     []int `json:"seat_id" binding:"required"`
	FilmID     int   `json:"film_id" binding:"required"`
}
