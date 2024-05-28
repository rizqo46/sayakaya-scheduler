package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"os"
	"sayakaya-scheduler/internal/notification"
	notificationConstant "sayakaya-scheduler/internal/notification/constant"
	"sayakaya-scheduler/internal/notification/dto"
	notificationRepository "sayakaya-scheduler/internal/notification/repository"
	"sayakaya-scheduler/internal/promos"
	promoDto "sayakaya-scheduler/internal/promos/dto"
	promosRepository "sayakaya-scheduler/internal/promos/repository"
	"sayakaya-scheduler/internal/users"
	userDto "sayakaya-scheduler/internal/users/dto"
	usersRepository "sayakaya-scheduler/internal/users/repository"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/vorobeyme/mailtrap-go/mailtrap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type schedulerApp struct {
	notificationService notification.NotificationService
	userService         users.UserService
	promoService        promos.PromoService
}

func (s *schedulerApp) processBirthdayPromo() {
	ctx := context.Background()
	// Get birthday users
	users, err := s.userService.FetchUsers(ctx, userDto.UserFilterField{
		Email:          "",
		VerifiedStatus: true,
		IsBirthday:     true,
	})
	if err != nil {
		log.Printf("failed to get users, err: %v", err)
		return
	}

	validUsersID := make([]int, 0, len(users))
	for i := range users {
		validUsersID = append(validUsersID, users[i].ID)
	}

	dateToday := time.Now().Round(0)

	// Create promo codes
	promo, err := s.promoService.GeneratePromoCode(ctx, promoDto.CreatePromoField{
		Name:         "Birthday Promo",
		StartDate:    dateToday,
		EndDate:      dateToday,
		Amount:       10000,
		ValidUsersID: validUsersID,
	})
	if err != nil {
		log.Printf("failed to create promo, err: %v", err)
		return
	}

	notificationParams := make([]dto.NotificationParams, 0, len(users))
	for i := range users {
		notificationParams = append(notificationParams, dto.NotificationParams{
			NotificationType: notificationConstant.NotificationTypeEmail,
			Subject:          "Birthday Special Promo!!",
			Body:             s.generateBirthDayEmailBody(users[i].Name, promo.PromoCode, promo.Amount),
			Target:           users[i].Email,
		})
	}

	// notify users
	s.notificationService.SendNotifications(ctx, notificationParams)
}

func (s *schedulerApp) generateBirthDayEmailBody(userName, promoCode string, amount int) string {
	tmpl, _ := template.New("template").Parse(
		`Happy Birthday {{.userName}}<br>
		<br>
		Kami punya hadiah voucher promo nih buat kamu. <br>
		Gunakan kode promo {{.promoCode}} untuk belanja apa saja di SayaKaya dan dapatkan potongan {{.amount}}.
		<br>
		Have a nice day.`,
	)
	data := map[string]interface{}{
		"userName":  userName,
		"promoCode": promoCode,
		"amount":    amount,
	}
	var b bytes.Buffer
	_ = tmpl.Execute(&b, &data)
	return b.String()
}

func initApp() schedulerApp {
	mailtrapApiToken := os.Getenv("MAILTRAP_API_TOKEN")
	client, err := mailtrap.NewSendingClient(mailtrapApiToken)
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	notificationRepo := notificationRepository.NewNotificationRepo(client)

	notificationService := notification.NewNotificationService(notificationRepo)

	promoRepo := promosRepository.NewPromoRepo()

	promoService := promos.NewPromoService(db, promoRepo)

	userRepo := usersRepository.NewUserRepo()

	userService := users.NewUserService(db, userRepo)

	return schedulerApp{
		notificationService: notificationService,
		userService:         userService,
		promoService:        promoService,
	}
}

func main() {
	app := initApp()
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	scheduler := gocron.NewScheduler(loc)
	scheduler.Every(1).Day().At("00:01").Do(func() { app.processBirthdayPromo() })
	scheduler.StartBlocking()
}
