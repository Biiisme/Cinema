package router

import (
	"cinema/handler"

	"github.com/labstack/echo"
)

type API struct {
	Echo        *echo.Echo
	Userhandler handler.Userhandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.Userhandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.Userhandler.HandleSignUp)

	api.Echo.POST("/user/sign-up", api.Userhandler.HandleSignUp)
	api.Echo.POST("/user/sign-up", api.Userhandler.HandleSignUp)
	api.Echo.POST("/user/sign-up", api.Userhandler.HandleSignUp)
	api.Echo.POST("/user/sign-up", api.Userhandler.HandleSignUp)

}
