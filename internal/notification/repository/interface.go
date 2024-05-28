package repository

import (
	"context"
	"sayakaya-scheduler/internal/notification/dto"
)

type NotificationRepo interface {
	SendNotificationEmail(ctx context.Context, param dto.NotificationParams) error
	SendNotificationWhatsapp(ctx context.Context, param dto.NotificationParams) error
}
