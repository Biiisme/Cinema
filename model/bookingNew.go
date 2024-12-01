package model

type Booking struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `json:"userId"`
	ScheduleID uint   `json:"scheduleId"`
	SeatID     uint   `json:"seatId"`
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
