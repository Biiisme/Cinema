package handler

import (
	"cinema/model"
	"cinema/model/req"
	"cinema/repository"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type FilmHandler struct {
	Repo repository.FilmRepo
}

func NewFilmHandler(repo repository.FilmRepo) *FilmHandler {
	return &FilmHandler{Repo: repo}
}

// SaveFilm lưu phim mới

func (u *FilmHandler) HandleSaveFilm(c echo.Context) error {
	// Kiểm tra Content-Type
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, model.Response{
			StatusCode: http.StatusUnsupportedMediaType,
			Message:    "Unsupported Media Type",
			Data:       nil,
		})
	}

	req := req.ReqFilm{}
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
	FilmId, err := uuid.NewUUID()
	// Tạo đối tượng Film từ yêu cầu
	film := model.Film{
		FilmId:      FilmId.String(),
		FilmName:    req.FilmName,
		Thoiluong:   req.Thoiluong,
		Gioihantuoi: req.Gioihantuoi,
		ImageFilm:   req.ImageFilm,
	}

	// Gọi hàm để lưu phim vào cơ sở dữ liệu
	film, err = u.FilmRepo.SaveFilm(c.Request().Context(), film)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Lưu phim thành công",
		Data:       film,
	})
}

// GetFilmByID lấy thông tin phim theo ID
func (h *FilmHandler) GetFilmByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	film, err := h.Repo.GetFilmByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(film)
}

// GetAllFilms lấy tất cả phim
func (h *FilmHandler) GetAllFilms(w http.ResponseWriter, r *http.Request) {
	films, err := h.Repo.GetAllFilms(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(films)
}
