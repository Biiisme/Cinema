package model

type Film struct {
	FilmId      string  `json:"-" db:"film_id, omitempty"`
	FilmName    string  `json:"filmName,omitempty" db:"film_name, omitempty"`
	Thoiluong   float32 `json:"timefull,omitempty" db:"timefull, omitempty"`
	Gioihantuoi int     `json:"limitAge,omitempty" db:"limitAge, omitempty"`
	ImageFilm   string  `json:"image,omitempty" db:"image, omitempty"`
}
