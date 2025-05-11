package model

import "time"

type User struct {
	UserId      string     `gorm:"primaryKey" json:"user_id"`
	Email       string     `json:"email,omitempty" gorm:"column:email"`
	Password    string     `json:"password,omitempty" gorm:"column:password"`
	PhoneNumber string     `json:"phone_number,omitempty" gorm:"column:phone_number"`
	BirthDate   *time.Time `json:"birth_date,omitempty" gorm:"column:birth_date"`
	Role        string     `json:"role,omitempty"`
	FullName    string     `json:"fullName,omitempty" gorm:"column:full_name"`
	Token       string     `json:"-,omitempty" gorm:"column:token"` // Thêm cột token vào đây
	CreatedAt   time.Time  `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt   time.Time  `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}
