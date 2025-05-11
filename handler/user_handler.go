package handler

import (
	"cinema/model"
	req "cinema/model/req"
	"cinema/repository"
	"fmt"

	"cinema/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

// HandleSignUp godoc
// @Summary      Đăng ký tài khoản mới
// @Description  API cho phép người dùng đăng ký tài khoản mới
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request  body      req.ReqSignUp  true  "Thông tin đăng ký"
// @Success      200      {object}  model.Response{data=model.User}
// @Failure      400      {object}  model.Response
// @Failure      409      {object}  model.Response
// @Failure      500      {object}  model.Response
// @Router       /user/sign-up [post]
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

	// Hash password
	hash := security.HashAndSalt([]byte(req.Password))

	//send email
	otp, err := security.GeneratorOTP(6)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}
	hashotp, _ := security.HashOTP(otp)
	senderr := security.SendSecretCodeToEmail(req.Email, otp, hashotp)
	if senderr != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "send email fail",
			Data:       nil,
		})
		return
	}
	//hash ID user
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

// HandleSignIn godoc
// @Summary      Đăng nhập
// @Description  API cho phép người dùng đăng nhập và nhận token
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request  body      req.ReqSignIn  true  "Thông tin đăng nhập"
// @Success      200      {object}  model.Response{data=map[string]interface{}}
// @Failure      400      {object}  model.Response
// @Failure      401      {object}  model.Response
// @Failure      409      {object}  model.Response
// @Failure      500      {object}  model.Response
// @Router       /user/sign-in [post]
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

	// Check users in the database
	user, err := u.UserRepo.CheckLogin(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Account does not exist",
			Data:       nil,
		})
		return
	}

	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Password is incorrect",
			Data:       nil,
		})
		return
	}

	token, err := security.GenerateJWT(user.UserId, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to create tokens",
			Data:       nil,
		})
		return
	}
	ip := c.ClientIP()
	fmt.Println(ip)
	if err := u.UserRepo.SaveToken(c.Request.Context(), user.UserId, token, ip); err != nil {
		c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    "Token unlock",
			Data:       token,
		})
		return
	}
	// Log in successfully
	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Login successfully",
		Data: gin.H{
			"user":  user,
			"token": token,
		},
	})
}

func (u *UserHandler) HandleGetUser(c *gin.Context) {
	userID, _ := c.Get("user_id")
	userIDUstring, _ := userID.(string)

	data, err := u.UserRepo.GetUser(c, userIDUstring)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusConflict,
			Message:    "Get user bug",
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Profile successfully",
		Data: gin.H{
			"data": data,
		},
	})
}

func (u *UserHandler) HandleUpdateUser(c *gin.Context) {
	// Kiểm tra Content-Type
	if c.GetHeader("Content-Type") != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, model.Response{
			StatusCode: http.StatusUnsupportedMediaType,
			Message:    "Unsupported Media Type",
			Data:       nil,
		})
		return
	}
	var req req.ReqUpdateProfile
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid data",
			Data:       validationErrors,
		})
		return
	}

	userID, _ := c.GetQuery("id")

	savedUser, err := u.UserRepo.UpdateUser(req, userID)
	if err != nil {
		c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Cập nhật user thành công",
		Data:       savedUser,
	})
}
