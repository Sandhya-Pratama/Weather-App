package router

import (
	"github.com/Sandhya-Pratama/weather-app/internal/http/handler"
	"github.com/labstack/echo/v4"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func PublicRoutes(userHandler *handler.UserHandler) []*Route  {	
	return []*Route{
		{
			Method: echo.GET,
			Path: "/users",
			Handler: userHandler.GetAllUsers,
		},
		{
			Method: echo.POST,
			Path: "/users",
			Handler: userHandler.CreateUser,
		},
	}
}

func PrivateRoutes() []*Route  {	
	return []*Route{
		{

		},
	}
}