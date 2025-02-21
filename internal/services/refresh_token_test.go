package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/models"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRefreshTokenService_RefreshToken(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	userMockRepo := NewMockIUserRepository(ctrlMock)

	type args struct {
		ctx          context.Context
		refreshToken string
		tokenClaim   helpers.ClaimToken
	}
	tests := []struct {
		name    string
		args    args
		want    models.RefreshTokenResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			want: models.RefreshTokenResponse{
				Token: gomock.Any().String(),
			},
			args: args{
				ctx:        context.Background(),
				tokenClaim: helpers.ClaimToken{},
			},
			wantErr: false,
			mockFn: func(args args) {
				userMockRepo.EXPECT().UpdateTokenByRefreshToken(args.ctx, gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "error update token",
			want: models.RefreshTokenResponse{
				Token: gomock.Any().String(),
			},
			args: args{
				ctx:        context.Background(),
				tokenClaim: helpers.ClaimToken{},
			},
			wantErr: true,
			mockFn: func(args args) {
				userMockRepo.EXPECT().UpdateTokenByRefreshToken(args.ctx, gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &RefreshTokenService{
				UserRepo: userMockRepo,
			}
			got, err := s.RefreshToken(tt.args.ctx, tt.args.refreshToken, tt.args.tokenClaim)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenService.RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
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
