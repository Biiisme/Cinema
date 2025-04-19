package model

import "time"

// Film ánh xạ tới bảng films
type Film struct {
	ID          int        `json:"id"`                 // ID duy nhất của phim
	Title       string     `json:"title"`              // Tên phim
	PosterURL   string     `json:"poster_url"`         // Link poster
	TrailerURL  string     `json:"trailer_url"`        // Link trailer
	Description string     `json:"description"`        // Mô tả
	Duration    int        `json:"duration"`           // Thời lượng
	ReleaseDate time.Time  `json:"release_date"`       // Ngày khởi chiếu
	EndDate     *time.Time `json:"end_date,omitempty"` // Ngày kết thúc (có thể null)
	//Genre        []string   `gorm:"type:jsonb" json:"genre"`  // Thể loại (mảng chuỗi)
	Director string `json:"director"` // Đạo diễn
	//Actors       []string   `gorm:"type:jsonb" json:"actors"` // Diễn viên (mảng chuỗi)
	Rated        string    `json:"rated"`          // Phân loại độ tuổi
	IsNowShowing bool      `json:"is_now_showing"` // Đang chiếu?
	IsComingSoon bool      `json:"is_coming_soon"` // Sắp chiếu?
	RatingAvg    float64   `json:"rating_avg"`     // Điểm TB
	RatingCount  int       `json:"rating_count"`   // Số lượt đánh giá
	CreatedAt    time.Time `json:"created_at"`     // Ngày tạo
	UpdatedAt    time.Time `json:"updated_at"`     // Ngày cập nhật
}
