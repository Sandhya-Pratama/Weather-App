package service

import (
	"context"
	"errors"

	"github.com/Sandhya-Pratama/weather-app/entity"
)

type LoginUseCase interface {
	Login(ctx context.Context, email, password string) (*entity.User, error)
}

type LoginRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type LoginService struct {
	repository LoginRepository
}

func NewLoginService(repository LoginRepository) *LoginService {
	return &LoginService{
		repository: repository,
	}
}

func (s *LoginService) Login(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := s.repository.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User with that email is not found")
	}

	if user.Password != password {
		return nil, errors.New("Incorrect password")
	}

	return user, nil
}
