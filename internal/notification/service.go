package notification

import (
	"context"
	"sayakaya-scheduler/internal/notification/constant"
	"sayakaya-scheduler/internal/notification/dto"
	"sayakaya-scheduler/internal/notification/repository"
	"sync"
)

type NotificationService struct {
	notificationRepo repository.NotificationRepo
}

func NewNotificationService(notificationRepo repository.NotificationRepo) NotificationService {
	return NotificationService{notificationRepo: notificationRepo}
}
func (s *NotificationService) SendNotifications(ctx context.Context, params []dto.NotificationParams) {
	var wg sync.WaitGroup

	for _, param := range params {
		wg.Add(1)
		go func(wg *sync.WaitGroup, ctx context.Context, param dto.NotificationParams) {
			defer wg.Done()
			s.SendNotification(ctx, param)
		}(&wg, ctx, param)
	}

	wg.Wait()
}
func (s *NotificationService) SendNotification(ctx context.Context, param dto.NotificationParams) {
	switch param.NotificationType {
	case constant.NotificationTypeEmail:
		s.notificationRepo.SendNotificationEmail(ctx, param)
	case constant.NotificationTypeWhatsapp:
		s.notificationRepo.SendNotificationWhatsapp(ctx, param)
	}
}
