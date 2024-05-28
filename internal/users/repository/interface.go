package repository

import (
	"context"
	"sayakaya-scheduler/internal/users/dto"
	"sayakaya-scheduler/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers(ctx context.Context, db *gorm.DB, param dto.UserFilterField) ([]models.User, error)
}
