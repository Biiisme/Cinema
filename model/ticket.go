package model

import "time"

type Ticket struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      string    `json:"user_id"`
	ScheduleID  int32     `json:"schedule_id"`
	Seat        string    `json:"seat"`
	TotalPrice  int64     `json:"total_price"`
	Code        int32     `json:"code"`
	Status      string    `json:"status"` // "booked" or "cancelled"
	PaymentDate time.Time `json:"payment_date"`
}
