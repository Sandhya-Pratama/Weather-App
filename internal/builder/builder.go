package builder

import (
	"github.com/Sandhya-Pratama/weather-app/internal/config"
	"github.com/Sandhya-Pratama/weather-app/internal/http/handler"
	"github.com/Sandhya-Pratama/weather-app/internal/http/router"
	"github.com/Sandhya-Pratama/weather-app/internal/repository"
	"github.com/Sandhya-Pratama/weather-app/internal/service"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB, redisClient *redis.Client) []*router.Route {
	userRepository := repository.NewUserRepository(db, redisClient)

	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)

	authHandler := handler.NewAuthHandler(loginService, tokenService)

	return router.PublicRoutes(authHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB, redisClient *redis.Client) []*router.Route {
	userRepository := repository.NewUserRepository(db, redisClient)

	userService := service.NewUserService(userRepository)

	userHandler := handler.NewUserHandler(cfg, userService)

	return router.PrivateRoutes(userHandler)
}
