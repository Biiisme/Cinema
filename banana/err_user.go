package banana

import "errors"

var (
	UserConfilict = errors.New("Người dùng đã tồn tại")
	SignUpFail    = errors.New("Đăng kí thất bại")
	UserNotFound  = errors.New("Không tìm thấy người dùng này")
	FilmConflict  = errors.New("film already exists")
	SaveFilmFail  = errors.New("failed to save film")
	FilmNotFound  = errors.New("film not found")
)
