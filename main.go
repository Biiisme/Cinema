package main

import (
	"cinema/db"
	"cinema/handler"
	"cinema/repository/repo_impl"
	"cinema/router"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	r := gin.Default()

	//  CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
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
	}

	api.SetupRouter()
	r.Run(":3000")
}
