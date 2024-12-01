package router

import (
	"cinema/handler"
	"cinema/security"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router          *gin.Engine
	UserHandler     handler.UserHandler
	FilmHandler     handler.FilmHandler
	ScheduleHandler handler.ScheduleHandler
	BookingHandler  handler.BookingHandler
}

func (r *API) SetupRouter() {
	// Định nghĩa các route cho user
	r.Router.POST("/user/sign-in", r.UserHandler.HandleSignIn)                            //Lấy thông tin user
	r.Router.POST("/user/sign-up", r.UserHandler.HandleSignUp)                            //Lưu user
	r.Router.GET("/movies", r.FilmHandler.GetAllFilms)                                    //lấy tất cả phim
	r.Router.GET("/film/:id", r.FilmHandler.GetFilmByID)                                  // Lấy thông tin phim theo ID
	r.Router.POST("/schedules", r.ScheduleHandler.HandleSaveSchedule)                     //Lưu lịch chiếu
	r.Router.GET("/schedules/film/:filmId", r.ScheduleHandler.HandleGetSchedulesByFilmID) //lấy lịch chiếu theo id film

	// Các route bảo vệ
	api := r.Router.Group("/api")
	api.Use(security.JWTAuthMiddleware())
	// Route dành cho khách hàng
	api.POST("/bookings", r.BookingHandler.CreateBooking)

	// Route dành cho admin
	admin := api.Group("/admin")
	admin.Use(security.AdminOnlyMiddleware())
	admin.POST("/addmovie", r.FilmHandler.HandleSaveFilm) //Add film

}
