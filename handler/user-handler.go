package handler

import (
	"cinema/model"
	"cinema/model/req"
	"cinema/security"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type Userhandler struct {
}

func (u *Userhandler) HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user123": "Bii",
		"email":   "Thangfakerlq@gmail.com",
	})
}

func (u *Userhandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSingUp{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	UserID, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	type User struct {
		RyanEmail string `json:"email"`
		FullName  string `json:"name"`
		Age       int    `json:"age"`
	}

	user := User{
		RyanEmail: "Thangfakerlq@gmail.com",
		FullName:  "Bii",
		Age:       21,
	}

	return c.JSON(http.StatusOK, user)
}
