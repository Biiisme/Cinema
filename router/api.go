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
	api := r.Router.Group("/api")
	api.POST("/user/sign-in", r.UserHandler.HandleSignIn)                             //login
	api.POST("/user/sign-up", r.UserHandler.HandleSignUp)                             //registration
	api.GET("/films", r.FilmHandler.GetAllFilms)                                      //Get_all_films
	api.GET("/film/:id", r.FilmHandler.GetFilmByID)                                   //Get_film_id
	api.POST("/schedules", r.ScheduleHandler.HandleSaveSchedule)                      //Create schedule
	api.GET("/schedules/film/:film-id", r.ScheduleHandler.HandleGetSchedulesByFilmID) //Get_schedule_filmID
	api.Use(security.JWTAuthMiddleware())
	customer := r.Router.Group("/customer")

	// Route for customer
	customer.POST("/bookings", r.BookingHandler.CreateBooking) //Same bookling

	// Route for admin
	admin := api.Group("/admin")
	admin.Use(security.AdminOnlyMiddleware())
	admin.POST("/add-film", r.FilmHandler.HandleSaveFilm) //Add film

}
