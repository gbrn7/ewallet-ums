package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/models"
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTokenValidationService_TokenValidation(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	userMockRepo := NewMockIUserRepository(ctrlMock)
	now := time.Now()
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *helpers.ClaimToken
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success validate token",
			args: args{
				ctx:   context.Background(),
				token: "token",
			},
			want: &helpers.ClaimToken{
				UserID:   1,
				Username: "username",
				Fullname: "fullname",
				Email:    "test@gmail.com",
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    helpers.GetEnv("APP_NAME", ""),
					IssuedAt:  jwt.NewNumericDate(now),
					ExpiresAt: jwt.NewNumericDate(now.Add(helpers.MapTypeToken["token"])),
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				userMockRepo.EXPECT().GetUserSessionByToken(args.ctx, args.token).Return(models.UserSession{}, nil)
			},
		},
		{
			name: "error",
			args: args{
				ctx:   context.Background(),
				token: "token",
			},
			want: &helpers.ClaimToken{
				UserID:   1,
				Username: "username",
				Fullname: "fullname",
				Email:    "test@gmail.com",
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    helpers.GetEnv("APP_NAME", ""),
					IssuedAt:  jwt.NewNumericDate(now),
					ExpiresAt: jwt.NewNumericDate(now.Add(helpers.MapTypeToken["token"])),
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				userMockRepo.EXPECT().GetUserSessionByToken(args.ctx, args.token).Return(models.UserSession{}, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		token, err := helpers.GenerateToken(tt.args.ctx, tt.want.UserID, tt.want.Username, tt.want.Fullname, tt.want.Email, now, "token")
		if err != nil {
			t.Errorf("error generate token : %v", err)
		}
		tt.args.token = token

		tt.mockFn(tt.args)

		t.Run(tt.name, func(t *testing.T) {
			s := &TokenValidationService{
				UserRepo: userMockRepo,
			}
			got, err := s.TokenValidation(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenValidationService.TokenValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenValidationService.TokenValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}
