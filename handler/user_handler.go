package handler

import (
	"cinema/model"
	req "cinema/model/req"
	"cinema/repository"
	"cinema/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

// HandleSignUp - Handles user registration
func (u *UserHandler) HandleSignUp(c *gin.Context) {
	// Kiểm tra Content-Type
	if c.GetHeader("Content-Type") != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, model.Response{
			StatusCode: http.StatusUnsupportedMediaType,
			Message:    "Unsupported Media Type",
			Data:       nil,
		})
		return
	}

	// Bind JSON request
	var req req.ReqSignUp
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	// Validate input
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	// Hash password and create user ID
	hash := security.HashAndSalt([]byte(req.Password))
	userId, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to create user ID",
			Data:       nil,
		})
		return
	}

	// Create object User
	user := model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     model.MEMBER.String(),
		Token:    "",
	}

	//Save users to database
	user, err = u.UserRepo.SaveUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    "Email is already in using",
			Data:       nil,
		})
		return
	}

	//Returns successful results
	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Registered successfully",
		Data:       user,
	})
}

// HandleSignIn - Handle user login
func (u *UserHandler) HandleSignIn(c *gin.Context) {
	// Bind JSON request
	var req req.ReqSignIn
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	// Validate input
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	// Kiểm tra người dùng trong database
	user, err := u.UserRepo.CheckLogin(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Tài khoản không tồn tại",
			Data:       nil,
		})
		return
	}

	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Mật khẩu không chính xác",
			Data:       nil,
		})
		return
	}

	token, err := security.GenerateJWT(user.UserId, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Không thể tạo token",
			Data:       nil,
		})
		return
	}

	// Đăng nhập thành công
	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Đăng nhập thành công",
		Data: gin.H{
			"user":  user,
			"token": token,
		},
	})
}