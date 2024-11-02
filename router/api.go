package router

import (
	"cinema/handler"
	"cinema/security"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router      *gin.Engine
	UserHandler handler.UserHandler
	FilmHandler handler.FilmHandler
}

func (r *API) SetupRouter() {
	// Định nghĩa các route cho user
	r.Router.POST("/user/sign-in", r.UserHandler.HandleSignIn) //Lấy thông tin user
	r.Router.POST("/user/sign-up", r.UserHandler.HandleSignUp) //Lưu user

	// Các route bảo vệ
	api := r.Router.Group("/api")
	api.Use(security.JWTAuthMiddleware())
	// Định nghĩa các route cho phim

	// Route dành cho khách hàng
	api.GET("/movies", r.FilmHandler.GetAllFilms)   // cả admin và khách hàng đều có quyền truy cập
	api.GET("/film/:id", r.FilmHandler.GetFilmByID) // Lấy thông tin phim theo ID
	// Route dành cho admin
	admin := api.Group("/admin")
	admin.Use(security.AdminOnlyMiddleware())
	admin.POST("/movies", r.FilmHandler.HandleSaveFilm)
	admin.GET("/movies", r.FilmHandler.GetAllFilms)
	admin.GET("/film/:id", r.FilmHandler.GetFilmByID)

}
