package req

type ReqSignUp struct {
	FullName string `json:"full_name,omitempty" validate:"required"` // tags
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type ReqUpdateProfile struct {
	FullName    string `json:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty" `
	Email       string `json:"email" validate:"required"`
	BirthDate   string `json:"birth_date,omitempty" `
}
