package main

import (
	"cinema/db"
	"cinema/handler"
	"cinema/repository/repo_impl"
	"cinema/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	sql := &db.Sql{
		Host:     "127.0.0.1",
		Port:     5432,
		Usename:  "postgres",
		Password: "19022003",
		DBName:   "userlogin",
	}
	sql.Connect()
	defer sql.Close()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.BodyLimit("2M"))
	userHandler := handler.Userhandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		Userhandler: userHandler,
	}

	api.SetupRouter()
	e.Logger.Fatal(e.Start(":3000"))
}
