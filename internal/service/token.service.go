package service

import (
	"context"
	"time"

	"github.com/Sandhya-Pratama/weather-app/common"
	"github.com/Sandhya-Pratama/weather-app/entity"
	"github.com/Sandhya-Pratama/weather-app/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	GenerateAccessToken(ctx context.Context, user *entity.User) (string, error)
}

type TokenService struct {
	cfg *config.Config
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

func (s *TokenService) GenerateAccessToken(ctx context.Context, user *entity.User) (string, error) {
	expiredTime := time.Now().Local().Add(10 * time.Minute)
	claims := common.JwtCustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := token.SignedString([]byte(s.cfg.JWT.SecretKey))

	if err != nil {
		return "", err
	}

	return encodedToken, nil
}
