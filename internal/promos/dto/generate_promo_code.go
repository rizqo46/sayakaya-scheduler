package dto

import "time"

type CreatePromoField struct {
	Name         string
	StartDate    time.Time
	EndDate      time.Time
	Amount       int
	ValidUsersID []int
}
