package handler

import (
	"cinema/model"
	"cinema/model/req"
	"cinema/repository"
	"cinema/utils"
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

// HandleSaveFilm godoc
// @Summary      Thêm phim mới
// @Description  Chỉ admin mới có quyền thêm phim mới vào hệ thống
// @Tags         film
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        Authorization  header    string  true  "Token (Bearer {token})"
// @Param        film           body      req.FilmReq  true  "Thông tin phim cần lưu"
// @Success      200            {object}  model.Response{data=model.Film}
// @Failure      400            {object}  model.Response{data=[]string}
// @Failure      401            {object}  model.Response
// @Failure      415            {object}  model.Response
// @Failure      409            {object}  model.Response
// @Router       /admin/add-film [post]
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

	var req req.FilmReq
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
		Title:       req.Title,
		PosterURL:   req.PosterURL,
		Description: req.Description,
		TrailerURL:  req.TrailerURL,
		Duration:    req.Duration,
		ReleaseDate: req.ReleaseDate,
		EndDate:     req.EndDate,
		//Genre:        req.Genre,
		Director: req.Director,
		//Actors:       req.Actors,
		Rated:        req.Rated,
		IsNowShowing: req.IsNowShowing,
		IsComingSoon: req.IsComingSoon,
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

// ShowAccount godoc
// @Summary      Show an film
// @Description  get string by ID
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Film ID"
// @Success      200  {object}  model.Response{data=model.Film}
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /film/{id} [get]
// GetFilmByID lấy thông tin phim theo ID
func (h *FilmHandler) GetFilmByID(c *gin.Context) {
	id := c.Param("id")

	film, err := h.FilmRepo.GetFilmByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Film not ",
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

// GetAllFilms godoc
// @Summary      Lấy danh sách tất cả phim
// @Description  Trả về danh sách tất cả phim có trong hệ thống
// @Tags         film
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Response{data=[]model.Film}
// @Failure      500  {object}  model.Response
// @Router       /films [get]
// GetAllFilms lấy tất cả phim
func (h *FilmHandler) GetAllFilms(c *gin.Context) {
	lengthstr := c.Query("length")
	pagestr := c.Query("page")
	page, length := utils.Pagination(pagestr, lengthstr)
	offset := (page - 1) * length

	films, err := h.FilmRepo.GetAllFilms(c.Request.Context(), offset, length)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	totalPage := h.FilmRepo.TotalPage(model.Film{}, length)

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Lấy tất cả phim thành công",
		Data: gin.H{
			"data": films,
			"pagination": gin.H{
				"current_page": page,
				"page_size":    length,
				"total_pages":  totalPage,
			},
		},
	})
}

// ShowAccount godoc
// @Summary      delete a film
// @Description  Chỉ admin mới có quyền thêm phim mới vào hệ thống
// @Tags         film
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        Authorization  header    string  true  "Token (Bearer {token})"
// @Param        id   path      int  true  "Film ID"
// @Success      200  {object}  model.Film
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /admin/delete-film/{id} [delete]
func (h *FilmHandler) DeleteFilmByID(c *gin.Context) {
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
	if err := h.FilmRepo.Delete(film); err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Delete Film not found",
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xóa phim thành công",
		Data:       film,
	})
}

func (h *FilmHandler) HandleUpdateFilm(c *gin.Context) {
	// Kiểm tra Content-Type
	if c.GetHeader("Content-Type") != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, model.Response{
			StatusCode: http.StatusUnsupportedMediaType,
			Message:    "Unsupported Media Type",
			Data:       nil,
		})
		return
	}
	var req req.FilmReq
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
	id := c.Param("id")
	savedFilm, err := h.FilmRepo.UpdateFilm(req, id)
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
		Message:    "Cập nhật phim thành công",
		Data:       savedFilm,
	})

}
