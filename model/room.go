package model

type Room struct {
	RoomID   int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}
