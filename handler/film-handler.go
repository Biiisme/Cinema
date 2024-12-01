package handler

import (
	"cinema/model"
	"cinema/model/req"
	"cinema/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FilmHandler struct {
	FilmRepo repository.FilmRepo
}

func NewFilmHandler(repo repository.FilmRepo) *FilmHandler {
	return &FilmHandler{FilmRepo: repo}
}

// HandleSaveFilm lưu phim mới
func (h *FilmHandler) HandleSaveFilm(c *gin.Context) {
	// Kiểm tra Content-Type
	if c.GetHeader("Content-Type") != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, model.Response{
			StatusCode: http.StatusUnsupportedMediaType,
			Message:    "Unsupported Media Type",
			Data:       nil,
		})
		return
	}

	var req req.ReqFilm
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

	film := model.Film{
		//Hỏi lại a tuệ
		FilmName:  req.FilmName,
		TimeFull:  req.Thoiluong,
		LimitAge:  req.Gioihantuoi,
		ImageFilm: req.ImageFilm,
	}

	savedFilm, err := h.FilmRepo.SaveFilm(c.Request.Context(), film)
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
		Message:    "Lưu phim thành công",
		Data:       savedFilm,
	})

}

// GetFilmByID lấy thông tin phim theo ID
func (h *FilmHandler) GetFilmByID(c *gin.Context) {
	id := c.Param("id")

	film, err := h.FilmRepo.GetFilmByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Film not found",
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Lấy phim thành công",
		Data:       film,
	})
}

// GetAllFilms lấy tất cả phim
func (h *FilmHandler) GetAllFilms(c *gin.Context) {
	films, err := h.FilmRepo.GetAllFilms(c.Request.Context())
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
		Data:       films,
	})
}
