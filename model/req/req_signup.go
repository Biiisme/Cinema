package req

type ReqSingUp struct {
	FullName string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
