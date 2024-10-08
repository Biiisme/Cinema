package handler

import (
	"cinema/model"
	req "cinema/model/req"
	"cinema/repository"
	"cinema/security"
	"net/http"

	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type Userhandler struct {
	UserRepo repository.UserRepo
}

// user sign up------------------------------------------------------------------
func (u *Userhandler) HandleSignUp(c echo.Context) error {
	// Kiểm tra Content-Type
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, model.Response{
			StatusCode: http.StatusUnsupportedMediaType,
			Message:    "Unsupported Media Type",
			Data:       nil,
		})
	}
	req := req.ReqSignUp{}
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

	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user := model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	/*token, err := security.GenToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token
	*/
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       user,
	})
}

func (u *Userhandler) HandleSignIn(c echo.Context) error {
	req := req.ReqSignIn{}
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
	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	//check password
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Đăng nhập thất bại",
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Đăng nhập thành công",
		Data:       user,
	})

}
