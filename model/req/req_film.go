package req

type ReqFilm struct {
	FilmName    string  `json:"filmName"`
	Thoiluong   float32 `json:"timefull"`
	Gioihantuoi int     `json:"limitAge"`
	ImageFilm   string  `json:"image"`
}
