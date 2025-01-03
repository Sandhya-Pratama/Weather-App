package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Sandhya-Pratama/weather-app/entity"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

const (
	UserKey = "users:all"
)

type UserRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewUserRepository(db *gorm.DB, redisClient *redis.Client) *UserRepository {
	return &UserRepository{
		db:          db,
		redisClient: redisClient,
	}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	users := make([]*entity.User, 0)
	val, err := r.redisClient.Get(context.Background(), UserKey).Result()
	if err != nil {
		err := r.db.WithContext(ctx).Find(&users).Error // SELECT * FROM users
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(users)
		if err != nil {
			return nil, err
		}

		// Set the data in Redis with an expiration time (e.g., 1 hour)
		err = r.redisClient.Set(ctx, UserKey, val, time.Duration(1)*time.Minute).Err()
		if err != nil {
			return nil, err
		}
		return users, nil
	}

	err = json.Unmarshal([]byte(val), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	query := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID)
	if user.Name != "" {
		query = query.Update("name", user.Name)
	}
	if user.Password != "" {
		query = query.Update("password", user.Password)
	}
	if user.Role != "" {
		query = query.Update("role", user.Role)
	}
	if user.Email != "" {
		query = query.Update("email", user.Email)
	}
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user with that email not found")
	}
	return user, nil
}
