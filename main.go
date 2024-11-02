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

	// Thực hiện migrate
	gormDB.Migrate()

	r := gin.Default()

	// Cấu hình CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(gormDB.DB), // Chuyển sang dùng GORM
	}

	filmHandler := handler.FilmHandler{
		FilmRepo: repo_impl.NewFilmRepoImpl(gormDB.DB), // Chuyển sang dùng GORM
	}

	api := router.API{
		Router:      r,
		UserHandler: userHandler,
		FilmHandler: filmHandler,
	}

	api.SetupRouter()
	r.Run(":3000")
}
