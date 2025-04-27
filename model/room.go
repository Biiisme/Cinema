package model

type Room struct {
	RoomID    int    `gorm:"primaryKey" json:"id"`
	Room_name string `gorm:"not null" json:"room_name"`
	CinemasID int    `gorm:"not null" json:"cinemaID"`
}
