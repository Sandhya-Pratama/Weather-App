package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/Sandhya-Pratama/weather-app/internal/builder"
	"github.com/Sandhya-Pratama/weather-app/internal/config"
	"github.com/Sandhya-Pratama/weather-app/internal/http/binder"
	"github.com/Sandhya-Pratama/weather-app/internal/http/server"
	"github.com/Sandhya-Pratama/weather-app/internal/http/validator"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg, err := config.NewConfig(".env")
	cekError(err)

	splash()

	db, err := buildGormDB(cfg.Postgres)
	cekError(err)

	redisClient := buildRedis(cfg)

	publicRoutes := builder.BuildPublicRoutes(cfg, db, redisClient)
	privateRoutes := builder.BuildPrivateRoutes(cfg, db, redisClient)

	echoBinder := &echo.DefaultBinder{}
	formValidator := validator.NewFormValidator()
	customBinder := binder.NewBinder(echoBinder, formValidator)

	srv := server.NewServer(
		cfg,
		customBinder,
		publicRoutes,
		privateRoutes,
	)

	runServer(srv, cfg.Port)

	waitForShutdown(srv)

	// users := make([]*entity.User, 0)
	// if err := db.Find(&users).Error; err != nil{
	// 	cekError(err)
	// }
	// for _, v := range users {
	// 	fmt.Println(v)
	// }

}

func buildRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       0,
	})
	return client
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

func buildGormDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// ini buar ngecek gormnya
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func splash() {
	colorReset := "\033[32m"
	splashText := `__      __               __  .__                              _____ ____________________ 
/  \    /  \ ____ _____ _/  |_|  |__   ___________            /  _  \\______   \______   \
\   \/\/   // __ \\__  \\   __\  |  \_/ __ \_  __ \  ______  /  /_\  \|     ___/|     ___/
 \        /\  ___/ / __ \|  | |   Y  \  ___/|  | \/ /_____/ /    |    \    |    |    |    
  \__/\  /  \___  >____  /__| |___|  /\___  >__|            \____|__  /____|    |____|    
       \/       \/     \/          \/     \/                        \/                    `
	fmt.Println(colorReset, strings.TrimSpace(splashText))
}

func cekError(err error) {
	if err != nil {
		panic(err)
	}
}
