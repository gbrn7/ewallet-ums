package services

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLogoutService_Logout(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	userMockRepo := NewMockIUserRepository(ctrlMock)
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx:   context.Background(),
				token: "token",
			},
			wantErr: false,
			mockFn: func(args args) {
				userMockRepo.EXPECT().DeleteUserSession(args.ctx, args.token).Return(nil)
			},
		},
		{
			name: "error",
			args: args{
				ctx:   context.Background(),
				token: "token",
			},
			wantErr: true,
			mockFn: func(args args) {
				userMockRepo.EXPECT().DeleteUserSession(args.ctx, args.token).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			s := &LogoutService{
				UserRepo: userMockRepo,
			}
			if err := s.Logout(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("LogoutService.Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
