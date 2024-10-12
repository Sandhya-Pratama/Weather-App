package server

import (
	"github.com/Sandhya-Pratama/weather-app/internal/config"
	"github.com/Sandhya-Pratama/weather-app/internal/http/binder"
	"github.com/Sandhya-Pratama/weather-app/internal/http/router"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
}

func NewServer(
	cfg *config.Config,
	binder *binder.Binder,
	publicRoutes, privateRoutes []*router.Route) *Server {

	e := echo.New()
	e.HideBanner = true
	e.Binder = binder

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	v1 := e.Group("/api/v1")

	//bisa dengan e.Get/e.Post, tapi ini cleancodenya supaya tidak satu persatu
	for _, public := range publicRoutes {
		v1.Add(public.Method, public.Path, public.Handler)
	}

	for _, private := range privateRoutes {
		v1.Add(private.Method, private.Path, private.Handler, JWTProtected(cfg.JWT.SecretKey))
	}

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
	return &Server{e}
}

func JWTProtected(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secretKey),
	})
}
