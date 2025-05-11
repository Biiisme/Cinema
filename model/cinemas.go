package model

type Cinemas struct {
	ID             int    `gorm:"primaryKey"`
	Cinema_name    string `gorm:"not null" json:"cinema_name"`
	Cinema_address string `gorm:"not null" json:"cinema_address"`
}
