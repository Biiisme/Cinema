package req

import "time"

type FilmReq struct {
	Title        string     `json:"title" binding:"required"`           // Tên phim (bắt buộc)
	PosterURL    string     `json:"poster_url" binding:"required,url"`  // Link ảnh poster
	TrailerURL   string     `json:"trailer_url" binding:"required,url"` // Link trailer
	Description  string     `json:"description" binding:"required"`     // Mô tả nội dung
	Duration     int        `json:"duration" binding:"required,min=1"`  // Thời lượng (>=1 phút)
	ReleaseDate  time.Time  `json:"release_date" binding:"required"`    // Ngày khởi chiếu
	EndDate      *time.Time `json:"end_date"`                           // Ngày kết thúc chiếu (có thể null)
	Genre        []string   `json:"genre" binding:"required,min=1"`     // Thể loại phim
	Director     string     `json:"director" binding:"required"`        // Tên đạo diễn
	Actors       []string   `json:"actors" binding:"required,min=1"`    // Danh sách diễn viên
	Rated        string     `json:"rated" binding:"required"`           // Phân loại độ tuổi
	IsNowShowing bool       `json:"is_now_showing"`                     // Đang chiếu hay không
	IsComingSoon bool       `json:"is_coming_soon"`                     // Sắp chiếu hay không
}
