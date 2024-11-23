package req

type ReqBooking struct {
	ScheduleID int    `json:"scheduleId" binding:"required"`
	SeatID     int    `json:"seatId" binding:"required"`
	UserID     string `json:"userName" binding:"required"`
}
