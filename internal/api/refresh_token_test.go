package api

import (
	"context"
	"encoding/json"
	"ewallet-ums/cmd/middleware"
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRefreshTokenHandler_RefreshToken(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockIRefrehTokenService(ctrlMock)

	now := time.Now()

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		exptectedBody      helpers.Response
		wantErr            bool
	}{
		{
			name:               "success",
			wantErr:            false,
			expectedStatusCode: 200,
			exptectedBody: helpers.Response{
				Message: constants.SuccessMessage,
				Data: map[string]interface{}{
					"token": "token",
				},
			},
			mockFn: func() {
				mockSvc.EXPECT().RefreshToken(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.RefreshTokenResponse{
					Token: "token",
				}, nil)
			},
		},
		{
			name:               "error",
			wantErr:            true,
			expectedStatusCode: 500,
			exptectedBody: helpers.Response{
				Message: constants.ErrServerError,
			},
			mockFn: func() {
				mockSvc.EXPECT().RefreshToken(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.RefreshTokenResponse{}, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			r := gin.Default()
			refreshTokenAPI := &RefreshTokenHandler{
				RefreshTokenService: mockSvc,
			}

			userV1 := r.Group("/user/v1")
			userV1.PUT("/refresh-token", middleware.MiddlewareValidateAuth, refreshTokenAPI.RefreshToken)

			endPoint := "/user/v1/refresh-token"

			w := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodPut, endPoint, nil)
			assert.NoError(t, err)
			token, err := helpers.GenerateToken(context.Background(), 1, "username", "fullname", "email@gmail.com", now, "token")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			if !tt.wantErr {

				res := w.Result()
				defer res.Body.Close()

				response := helpers.Response{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.exptectedBody, response)
			}
		})
	}
}
