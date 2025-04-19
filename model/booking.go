package model

type Booking struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	UserID     string `json:"user_id"`
	FilmID     int    `json:"film_id"`
	ScheduleID int    `json:"schedule_id"`
	SeatID     int    `json:"seat_id"`
	Status     string `json:"status"` // "booked" or "cancelled"
}

type BookingInvoice struct {
	UserName    string   `json:"userName"`
	FilmName    string   `json:"filmName"`
	ShowDate    string   `json:"showDate"`
	ShowTime    string   `json:"showTime"`
	RoomName    string   `json:"roomName"`
	SeatNumbers []string `json:"seatNumbers"`
	TotalPrice  float64  `json:"totalPrice"`
}
