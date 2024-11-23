package model

import "time"

// Film ánh xạ tới bảng films
type Film struct {
	FilmID    string    `gorm:"primaryKey" json:"-"`             // Khoá chính
	FilmName  string    `json:"filmName,omitempty"`              // Tên phim
	TimeFull  float32   `json:"timefull,omitempty"`              // Thời lượng phim
	LimitAge  int       `json:"limitAge,omitempty"`              // Giới hạn tuổi
	ImageFilm string    `json:"image,omitempty"`                 // Đường dẫn ảnh
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"` // Thời gian tạo
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"` // Thời gian cập nhật
}
