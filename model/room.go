package model

type Room struct {
	ID        int     `gorm:"primaryKey" json:"id"`
	Room_name string  `gorm:"not null" json:"room_name"`
	CinemaID  int     `gorm:"not null" json:"cinema_id"`
	Cinema    Cinemas `gorm:"foreignKey:CinemaID;references:ID" json:"cinemas" `
}
