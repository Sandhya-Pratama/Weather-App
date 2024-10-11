package service

import (
	"context"

	"github.com/Sandhya-Pratama/weather-app/entity"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}
type UserService struct{
	repository UserRepository
}	

func NewUserService(repository UserRepository) *UserService {
    return &UserService{
        repository: repository,
    }
}

func (s *UserService) FindAll(ctx context.Context) ([]*entity.User, error) {
	return s.repository.FindAll(ctx)
}

func (s *UserService) Create(ctx context.Context, user *entity.User) error{
	return s.repository.Create(ctx, user)
}
