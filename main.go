package main

import (
	"cinema/db"
	"cinema/handler"
	"cinema/repository/repo_impl"
	"cinema/router"
	"cinema/utils"
	"log"
	"time"

	_ "cinema/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:3000
// @BasePath /api
func main() {
	gormDB := db.NewGormDB("127.0.0.1", "postgres", "19022003", "userlogin", 5432)
	defer func() {
		sqlDB, err := gormDB.DB.DB()
		if err != nil {
			log.Fatalf("Could not get DB: %v", err)
		}
		sqlDB.Close()
	}()

	//  migrate
	gormDB.Migrate()

	/* Init Cache */
	if err := utils.InitCache(); err != nil {
		return
	}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//  CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(gormDB.DB), //  GORM
	}

	filmHandler := handler.FilmHandler{
		FilmRepo: repo_impl.NewFilmRepoImpl(gormDB.DB), // GORM
	}
	cinemaHandler := handler.CinemaHandler{
		CinemaRepo: repo_impl.NewCinemaRepoImpl(gormDB.DB), // GORM
	}
	seatHandler := handler.SeatHandler{
		SeatRepo: repo_impl.NewSeatRepoImpl(gormDB.DB), // GORM
	}
	scheduleHandler := handler.ScheduleHandler{
		ScheduleRepo: *repo_impl.NewScheduleRepoImpl(gormDB.DB),
	}
	bookhingHandler := handler.BookingHandler{
		BookingRepo: repo_impl.NewBookingRepo(gormDB.DB),
	}
	api := router.API{
		Router:          r,
		UserHandler:     userHandler,
		FilmHandler:     filmHandler,
		ScheduleHandler: scheduleHandler,
		BookingHandler:  bookhingHandler,
		CinemaHandler:   cinemaHandler,
		SeatHandler:     seatHandler,
	}

	api.SetupRouter()

	r.Run(":3000")
}
