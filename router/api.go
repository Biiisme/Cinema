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
	CinemaHandler   handler.CinemaHandler
	SeatHandler     handler.SeatHandler
}

func (r *API) SetupRouter() {
	// Định nghĩa các route cho user
	api := r.Router.Group("/api")
	api.POST("/user/sign-in", r.UserHandler.HandleSignIn)                        //login
	api.POST("/user/sign-up", r.UserHandler.HandleSignUp)                        //registration
	api.GET("/films", r.FilmHandler.GetAllFilms)                                 //Get_all_films
	api.GET("/film/:id", r.FilmHandler.GetFilmByID)                              //Get_film_id
	api.POST("/schedules", r.ScheduleHandler.HandleSaveSchedule)                 //Create schedule
	api.GET("/schedules/film/:id", r.ScheduleHandler.HandleGetSchedulesByFilmID) //Get_schedule_filmID
	api.GET("/schedules/detail/:id", r.ScheduleHandler.HandleReadSchedule)
	api.GET("/cinemas", r.CinemaHandler.GetAllCinemas)
	api.GET("/seats/:id", r.SeatHandler.GetSeatbyFilmID)

	customer := r.Router.Group("/customer")
	customer.Use(security.JWTAuthMiddleware())
	// Route for customer
	customer.GET("/user/profile", r.UserHandler.HandleGetUser)
	customer.PATCH("/update-profile/:id", r.UserHandler.HandleUpdateUser)
	customer.POST("/hold-seat", r.BookingHandler.HoldSeat)
	customer.GET("/get-hold-seat", r.BookingHandler.GetHoldSeatInfo)
	//	customer.POST("/bookings", r.BookingHandler.CreateBooking) //Same bookling

	// Route for admin
	admin := api.Group("/admin")
	admin.Use(security.AdminOnlyMiddleware())
	admin.POST("/add-film", r.FilmHandler.HandleSaveFilm)          //Add film
	admin.DELETE("/delete-film/:id", r.FilmHandler.DeleteFilmByID) //Delete film
	admin.PATCH("/update-film/:id", r.FilmHandler.HandleUpdateFilm)
}
