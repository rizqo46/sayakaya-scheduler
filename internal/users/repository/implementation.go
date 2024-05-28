package repository

import (
	"context"
	"sayakaya-scheduler/internal/users/dto"
	"sayakaya-scheduler/models"
	"time"

	"gorm.io/gorm"
)

type userRepoImpl struct{}

func (u userRepoImpl) GetUsers(ctx context.Context, db *gorm.DB, param dto.UserFilterField) ([]models.User, error) {
	query := db.WithContext(ctx).Model(&models.User{})

	if param.Email != "" {
		query.Where("email = ?", param.Email)
	}

	if param.VerifiedStatus {
		query.Where("verified_status IS TRUE")
	}

	if param.IsBirthday {
		dateToday := time.Now().Format("01-02")
		query.Where("TO_CHAR(birthday, 'MM-DD') = ?", dateToday)
	}

	users := []models.User{}
	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func NewUserRepo() UserRepo {
	return userRepoImpl{}
}
