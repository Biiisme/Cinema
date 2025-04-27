package req

type RoomReq struct {
	Room_name string ` json:"room_name" binding:"required"`
	CinemasID int    `json:"cinemaID" binding:"required"`
}
