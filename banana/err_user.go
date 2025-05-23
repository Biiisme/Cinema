package banana

import "errors"

var (
	UserConfilict       = errors.New("Người dùng đã tồn tại")
	SignUpFail          = errors.New("Đăng kí thất bại")
	UserNotFound        = errors.New("Không tìm thấy người dùng này")
	SaveUserFail        = errors.New("failed to save user")
	FilmConflict        = errors.New("film already exists")
	SaveFilmFail        = errors.New("failed to save film")
	FilmNotFound        = errors.New("film not found")
	SaveScheduleFail    = errors.New("failed to save film")
	SaveBooking         = errors.New("failed to save film")
	ForeignKeyViolation = errors.New("Foreign key constraint violated")
)
