package model

type Booking struct {
	ID         int    `gorm:"primaryKey"`
	UserID     string `json:"userId"`
	ScheduleID int    `json:"scheduleId"`
	SeatID     int    `json:"seatId"`
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
