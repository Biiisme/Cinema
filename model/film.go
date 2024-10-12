package model

import "time"

// Film ánh xạ tới bảng films
type Film struct {
	FilmId      string    `gorm:"primaryKey" json:"-"`             // Khoá chính
	FilmName    string    `json:"filmName,omitempty"`              // Tên phim
	Thoiluong   float32   `json:"timefull,omitempty"`              // Thời lượng phim
	Gioihantuoi int       `json:"limitAge,omitempty"`              // Giới hạn tuổi
	ImageFilm   string    `json:"image,omitempty"`                 // Đường dẫn ảnh
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"` // Thời gian tạo
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"` // Thời gian cập nhật
}
