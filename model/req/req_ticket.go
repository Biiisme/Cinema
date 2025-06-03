package req

import "time"

type HoldSeatRequest struct {
	ScheduleId int        `json:"schedule_id" binding:"required"`
	Seats      []int      `json:"seats" binding:"required"`
	SeatName   []string   `json:"seat_name" binding:"required"`
	CreatedAt  *time.Time `json:"create_at"`
}
type BookingTicketRequest struct {
	BookingSessionID string `json:"bookingSessionId"`
}
