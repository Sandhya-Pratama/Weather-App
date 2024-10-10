package builder

import (
	"github.com/Sandhya-Pratama/weather-app/internal/config"
	"github.com/Sandhya-Pratama/weather-app/internal/http/router"
	"gorm.io/gorm"
)

type Builder struct {
}

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {

	return router.PublicRoutes()
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route{
	
	return router.PublicRoutes()
}