package repository

import (
	"context"
	"sayakaya-scheduler/models"

	"gorm.io/gorm"
)

type promoRepoImpl struct{}

func (p promoRepoImpl) CreatePromo(ctx context.Context, db *gorm.DB, model *models.Promo) error {
	return db.WithContext(ctx).Create(model).Error
}

func (p promoRepoImpl) CreatePromoEligiblities(ctx context.Context, db *gorm.DB, model []models.PromoEligibility) error {
	return db.WithContext(ctx).Create(&model).Error
}

func NewPromoRepo() PromoRepo {
	return promoRepoImpl{}
}
