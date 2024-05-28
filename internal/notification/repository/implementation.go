package repository

import (
	"context"
	"os"
	"sayakaya-scheduler/internal/notification/dto"

	"github.com/vorobeyme/mailtrap-go/mailtrap"
)

type notificationRepoImpl struct {
	mailer *mailtrap.SendingClient
}

func NewNotificationRepo(mailer *mailtrap.SendingClient) NotificationRepo {
	return &notificationRepoImpl{mailer: mailer}
}

func (r *notificationRepoImpl) SendNotificationEmail(ctx context.Context, param dto.NotificationParams) error {
	sendEmailReq := &mailtrap.SendEmailRequest{
		From: mailtrap.EmailAddress{
			Email: os.Getenv("EMAIL_SENDER"),
		},
		To: []mailtrap.EmailAddress{
			{
				Email: param.Target,
			},
		},
		Subject: param.Subject,
		HTML:    param.Body,
	}
	_, _, err := r.mailer.Send(sendEmailReq)
	return err
}

func (r *notificationRepoImpl) SendNotificationWhatsapp(ctx context.Context, param dto.NotificationParams) error {
	// dummy function
	return nil
}
