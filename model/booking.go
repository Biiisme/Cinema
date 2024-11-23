package model

import "time"

type Booking struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	ScheduleID int       `json:"schedule_id"`
	SeatID     int       `json:"seat_id"`
	UserName   string    `json:"user_name"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
