package server

import (
	"net/http"

	"github.com/Sandhya-Pratama/weather-app/internal/config"
	"github.com/Sandhya-Pratama/weather-app/internal/http/binder"
	"github.com/Sandhya-Pratama/weather-app/internal/http/router"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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
		session.Middleware(sessions.NewCookieStore([]byte(cfg.Session.SecretKey))),
	)

	v1 := e.Group("/api/v1")

	//bisa dengan e.Get/e.Post, tapi ini cleancodenya supaya tidak satu persatu
	for _, public := range publicRoutes {
		v1.Add(public.Method, public.Path, public.Handler)
	}

	for _, private := range privateRoutes {
		v1.Add(private.Method, private.Path, private.Handler, JWTProtected(cfg.JWT.SecretKey), RBACmiddleware(private.Roles...))
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

// SessionProtected middleware to protect endpoints with session
func SessionProtected() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			sess, _ := session.Get("auth-sessions", ctx)
			if sess.Values["token"] == nil {
				return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "silahkan login terlebih dahulu"})
			}
			ctx.Set("user", sess.Values["token"])
			return next(ctx)
		}
	}
}

func RBACmiddleware(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get("auth-sessions", c)
			role := sess.Values["role"].(string)
			for _, v := range roles {
				if v == role {
					return next(c)
				} else {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
				}
			}
			return next(c)
		}
	}
}
