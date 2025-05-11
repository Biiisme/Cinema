package handler

import (
	"cinema/model"
	"cinema/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CinemaHandler struct {
	CinemaRepo repository.CinemaRepo
}

func NewCinemaHandler(repo repository.CinemaRepo) *CinemaHandler {
	return &CinemaHandler{CinemaRepo: repo}
}

func (u *CinemaHandler) GetAllCinemas(c *gin.Context) {
	cinemas, err := u.CinemaRepo.GetAllCinemas(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Lấy tất cả phim thành công",
		Data:       cinemas,
	})
}
