package users

import (
	"context"
	"sayakaya-scheduler/internal/users/dto"
	"sayakaya-scheduler/internal/users/repository"
	"sayakaya-scheduler/models"

	"gorm.io/gorm"
)

type UserService struct {
	db       *gorm.DB
	userRepo repository.UserRepo
}

func NewUserService(db *gorm.DB, userRepo repository.UserRepo) UserService {
	return UserService{db: db, userRepo: userRepo}
}

func (s *UserService) FetchUsers(ctx context.Context, param dto.UserFilterField) ([]models.User, error) {
	return s.userRepo.GetUsers(ctx, s.db, param)
}
