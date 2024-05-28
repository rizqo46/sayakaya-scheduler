package repository

import (
	"context"
	"sayakaya-scheduler/models"

	"gorm.io/gorm"
)

type PromoRepo interface {
	CreatePromo(ctx context.Context, db *gorm.DB, model *models.Promo) error
	CreatePromoEligiblities(ctx context.Context, db *gorm.DB, model []models.PromoEligibility) error
}
