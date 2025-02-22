package api

import (
	context "context"
	"encoding/json"
	"ewallet-ums/cmd/middleware"
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLogoutHandler_Logout(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockILogoutService(ctrlMock)

	tests := []struct {
		name               string
		expectedStatusCode int
		exptectedBody      helpers.Response
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: 200,
			exptectedBody: helpers.Response{
				Message: constants.SuccessMessage,
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().Logout(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:               "error",
			expectedStatusCode: 500,
			exptectedBody: helpers.Response{
				Message: constants.ErrServerError,
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().Logout(gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			r := gin.Default()

			api := &LogoutHandler{
				LogoutService: mockSvc,
			}

			userV1 := r.Group("/user/v1")
			userV1.DELETE("/logout", middleware.MiddlewareValidateAuth, api.Logout)

			endPoint := "/user/v1/logout"
			w := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodDelete, endPoint, nil)
			assert.NoError(t, err)
			token, err := helpers.GenerateToken(context.Background(), 1, "username", "fullname", "email@gmail.com", time.Now(), "token")
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
