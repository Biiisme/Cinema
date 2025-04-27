package model

import "time"

// Schedule ánh xạ tới bảng schedules
type Schedule struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	FilmID   int       `gorm:"not null" json:"filmId"`   // ID phim
	RoomID   int       `gorm:"not null" json:"roomId"`   // ID phòng chiếu
	CinemaID int       `gorm:"not null" json:"cinemaId"` // ID Rạp chiếu
	ShowDate time.Time `gorm:"not null" json:"showDate"` // Ngày chiếu
	ShowTime time.Time `gorm:"not null" json:"showTime"` // Giờ chiếu

}
