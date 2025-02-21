package services

import (
	"context"
	"ewallet-ums/internal/models"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginService_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	userMockRepo := NewMockIUserRepository(ctrlMock)

	password, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		t.Errorf("Error to generate password: %v", err)
		return
	}

	type args struct {
		ctx context.Context
		req models.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		want    models.LoginResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: models.LoginRequest{
					Username: "username",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				userMockRepo.EXPECT().GetUserByUsername(args.ctx, args.req.Username).Return(models.User{
					ID:          1,
					Username:    "username",
					Email:       "email@gmail.com",
					PhoneNumber: "phone_number",
					Fullname:    "fullname",
					Address:     "address",
					Dob:         "dob",
					Password:    string(password),
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil)
				userMockRepo.EXPECT().InsertNewUserSession(args.ctx, gomock.Any()).Return(nil)
			},
		},
		{
			name: "error get user",
			args: args{
				ctx: context.Background(),
				req: models.LoginRequest{
					Username: "username",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				userMockRepo.EXPECT().GetUserByUsername(args.ctx, args.req.Username).Return(models.User{}, assert.AnError)
			},
		},
		{
			name: "error insert user sesstion",
			args: args{
				ctx: context.Background(),
				req: models.LoginRequest{
					Username: "username",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				userMockRepo.EXPECT().GetUserByUsername(args.ctx, args.req.Username).Return(models.User{
					ID:          1,
					Username:    "username",
					Email:       "email@gmail.com",
					PhoneNumber: "phone_number",
					Fullname:    "fullname",
					Address:     "address",
					Dob:         "dob",
					Password:    string(password),
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}, nil)
				userMockRepo.EXPECT().InsertNewUserSession(args.ctx, gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &LoginService{
				UserRepo: userMockRepo,
			}
			got, err := s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
