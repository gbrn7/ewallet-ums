package api

import (
	"context"
	"ewallet-ums/cmd/proto/tokenvalidation"
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	reflect "reflect"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTokenValidationHandler_ValidateToken(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockITokenValidationService(ctrlMock)

	tests := []struct {
		name    string
		want    *tokenvalidation.TokenResponse
		wantErr bool
		mockFn  func()
	}{
		{
			name: "success",
			want: &tokenvalidation.TokenResponse{
				Message: constants.SuccessMessage,
				Data: &tokenvalidation.UserData{
					UserId:   1,
					Username: "username",
					FullName: "fullname",
					Email:    "email@gmail.com",
				},
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().TokenValidation(gomock.Any(), gomock.Any()).Return(&helpers.ClaimToken{
					UserID:   1,
					Username: "username",
					Fullname: "fullname",
					Email:    "email@gmail.com",
				}, nil)
			},
		},
		{
			name: "error",
			want: &tokenvalidation.TokenResponse{
				Message: constants.ErrServerError,
			},
			wantErr: true,
			mockFn: func() {
				mockSvc.EXPECT().TokenValidation(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			s := &TokenValidationHandler{
				TokenValidationService: mockSvc,
			}

			token, err := helpers.GenerateToken(context.Background(), 1, "username", "fullname", "email@gmail.com", time.Now(), "token")
			assert.NoError(t, err)

			got, err := s.ValidateToken(context.Background(), &tokenvalidation.TokenRequest{
				Token: token,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
