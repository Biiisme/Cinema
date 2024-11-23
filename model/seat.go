package model

type Seat struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	RoomID     int    `json:"room_id"`
	SeatNumber string `json:"seat_number"`
	IsReserved bool   `json:"is_reserved"`
}
