package router

import (
	"cinema/handler"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router      *gin.Engine
	UserHandler handler.UserHandler
	FilmHandler handler.FilmHandler
}

func (api *API) SetupRouter() {
	// Định nghĩa các route cho user
	api.Router.POST("/user/sign-in", api.UserHandler.HandleSignIn) //Lấy thông tin user
	api.Router.POST("/user/sign-up", api.UserHandler.HandleSignUp) //Lưu user

	// Định nghĩa các route cho phim
	api.Router.POST("/film", api.FilmHandler.HandleSaveFilm) // Lưu phim mới
	api.Router.GET("/film/:id", api.FilmHandler.GetFilmByID) // Lấy thông tin phim theo ID
	api.Router.GET("/films", api.FilmHandler.GetAllFilms)    // Lấy tất cả phim

}
