package api

import (
	"bytes"
	"encoding/json"
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockILoginService(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		exptectedBody      helpers.Response
		wantErr            bool
	}{
		{
			name:               "success",
			expectedStatusCode: 200,
			wantErr:            false,
			exptectedBody: helpers.Response{
				Message: constants.SuccessMessage,
				Data: map[string]interface{}{
					"user_id":       float64(1),
					"username":      "username",
					"full_name":     "fullname",
					"email":         "email",
					"token":         "token",
					"refresh_token": "refresh_token",
				},
			},
			mockFn: func() {
				mockSvc.EXPECT().Login(gomock.Any(), models.LoginRequest{
					Username: "username",
					Password: "password",
				}).Return(models.LoginResponse{
					UserID:       1,
					Username:     "username",
					Fullname:     "fullname",
					Email:        "email",
					Token:        "token",
					RefreshToken: "refresh_token",
				}, nil)
			},
		},
		{
			name:               "error",
			expectedStatusCode: 500,
			wantErr:            true,
			exptectedBody: helpers.Response{
				Message: constants.SuccessMessage,
			},
			mockFn: func() {
				mockSvc.EXPECT().Login(gomock.Any(), models.LoginRequest{
					Username: "username",
					Password: "password",
				}).Return(models.LoginResponse{}, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			r := gin.Default()
			api := &LoginHandler{
				LoginService: mockSvc,
			}

			userV1 := r.Group("/user/v1")
			userV1.POST("/login", api.Login)
			endPoint := "/user/v1/login"

			w := httptest.NewRecorder()
			model := models.LoginRequest{
				Username: "username",
				Password: "password",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(val)

			req, err := http.NewRequest(http.MethodPost, endPoint, body)
			assert.NoError(t, err)

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
