package main

import (
	"sayakaya-scheduler/internal/notification"
	"sayakaya-scheduler/internal/promos"
	"sayakaya-scheduler/internal/users"
	"testing"
)

func Test_schedulerApp_generateBirthDayEmailBody(t *testing.T) {
	type fields struct {
		notificationService notification.NotificationService
		userService         users.UserService
		promoService        promos.PromoService
	}
	type args struct {
		userName  string
		promoCode string
		amount    int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "",
			fields: fields{},
			args: args{
				userName:  "Abyan",
				promoCode: "19872h77",
				amount:    1000,
			},
			want: "Happy Birthday Abyan<br>\n\t\t<br>\n\t\tKami punya hadiah voucher promo nih buat kamu. <br>\n\t\tGunakan kode promo 19872h77 untuk belanja apa saja di SayaKaya dan dapatkan potongan 1000.\n\t\t<br>\n\t\tHave a nice day.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &schedulerApp{
				notificationService: tt.fields.notificationService,
				userService:         tt.fields.userService,
				promoService:        tt.fields.promoService,
			}
			if got := s.generateBirthDayEmailBody(tt.args.userName, tt.args.promoCode, tt.args.amount); got != tt.want {
				t.Errorf("schedulerApp.generateBirthDayEmailBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
