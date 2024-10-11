package builder

import (
	"github.com/Sandhya-Pratama/weather-app/internal/config"
	"github.com/Sandhya-Pratama/weather-app/internal/http/handler"
	"github.com/Sandhya-Pratama/weather-app/internal/http/router"
	"github.com/Sandhya-Pratama/weather-app/internal/repository"
	"github.com/Sandhya-Pratama/weather-app/internal/service"
	"gorm.io/gorm"
)

type Builder struct {
}

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	return router.PublicRoutes(userHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {

	return router.PrivateRoutes()
}
