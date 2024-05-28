package models

import "time"

type Promo struct {
	ID        int
	Name      string
	Amount    int
	PromoCode string
	StartDate time.Time `gorm:"type:date"`
	EndDate   time.Time `gorm:"type:date"`
}

func (Promo) TableName() string {
	return "promos"
}

type PromoEligibility struct {
	PromoID   int       `gorm:"primaryKey"`
	UserID    int       `gorm:"primaryKey"`
	ClaimedAt time.Time `gorm:"type:timestamptz"`
}

func (PromoEligibility) TableName() string {
	return "promo_eligibilities"
}
