package model

type Seat struct {
	ID         int     `gorm:"primaryKey" json:"id"`
	RoomID     int     `gorm:"not null" json:"roomId"`                      // ID phòng
	SeatNumber string  `gorm:"not null" json:"seatNumber"`                  // Số ghế
	SeatType   string  `gorm:"not null" json:"seatType"`                    // Loại ghế (normal hoặc couple)
	Price      float64 `gorm:"not null" json:"price"`                       // Giá ghế
	Room       Room    `gorm:"foreignKey:RoomID;references:ID" json:"room"` // Quan hệ với bảng Room
}
