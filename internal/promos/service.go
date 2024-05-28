package promos

import (
	"context"
	"log"
	"math/rand"
	"sayakaya-scheduler/internal/promos/dto"
	"sayakaya-scheduler/internal/promos/repository"
	"sayakaya-scheduler/models"

	"gorm.io/gorm"
)

type PromoService struct {
	db        *gorm.DB
	promoRepo repository.PromoRepo
}

func NewPromoService(db *gorm.DB, promoRepo repository.PromoRepo) PromoService {
	return PromoService{db: db, promoRepo: promoRepo}
}

func createPromoCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := make([]byte, 8)
	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}
	return string(id)
}

func (s *PromoService) GeneratePromoCode(ctx context.Context, param dto.CreatePromoField) (*models.Promo, error) {
	promo := models.Promo{
		Name:      param.Name,
		PromoCode: createPromoCode(),
		Amount:    param.Amount,
		StartDate: param.StartDate,
		EndDate:   param.EndDate,
	}

	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Rollback()

	err := s.promoRepo.CreatePromo(ctx, tx, &promo)
	if err != nil {
		log.Printf("failed to create promo, err: %v", err)
		return nil, err
	}

	promoEligibilities := make([]models.PromoEligibility, 0, len(param.ValidUsersID))
	for i := range param.ValidUsersID {
		promoEligibilities = append(promoEligibilities, models.PromoEligibility{
			PromoID: promo.ID,
			UserID:  param.ValidUsersID[i],
		})
	}

	err = s.promoRepo.CreatePromoEligiblities(ctx, tx, promoEligibilities)
	if err != nil {
		log.Printf("failed to create promo eligibilities, err: %v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return &promo, nil
}
