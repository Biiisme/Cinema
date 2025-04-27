package req

type ReqSchedule struct {
	FilmID   int    `json:"filmId" binding:"required"`
	RoomID   int    `json:"roomId" binding:"required"`
	CinemaID int    `json:"cinemaID" binding:"required"`
	ShowDate string `json:"showDate" binding:"required"`
	ShowTime string `json:"showTime" binding:"required"`
}
