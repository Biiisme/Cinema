package model

type Room struct {
	RoomID int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"not null" json:"name"`
}
