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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler_Register(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockIRegisterService(ctrlMock)

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
			expectedStatusCode: 200,
			wantErr:            false,
			exptectedBody: helpers.Response{
				Message: constants.SuccessMessage,
				Data: map[string]interface{}{
					"id":           float64(1),
					"username":     "username",
					"email":        "email@gmail.com",
					"phone_number": "phone_number",
					"full_name":    "fullname",
					"address":      "address",
					"dob":          "dob",
					"password":     "password",
					"created_at":   now.Format(time.RFC3339Nano),
					"updated_at":   now.Format(time.RFC3339Nano),
				},
			},
			mockFn: func() {
				mockSvc.EXPECT().Register(gomock.Any(), models.User{
					Username:    "username",
					Email:       "email@gmail.com",
					PhoneNumber: "phone_number",
					Fullname:    "fullname",
					Address:     "address",
					Dob:         "dob",
					Password:    "password",
				}).Return(models.User{
					ID:          1,
					Username:    "username",
					Email:       "email@gmail.com",
					PhoneNumber: "phone_number",
					Fullname:    "fullname",
					Address:     "address",
					Dob:         "dob",
					Password:    "password",
					CreatedAt:   now,
					UpdatedAt:   now,
				}, nil)
			},
		},
		{
			name:               "error",
			wantErr:            true,
			expectedStatusCode: http.StatusInternalServerError,
			exptectedBody: helpers.Response{
				Message: constants.ErrServerError,
			},
			mockFn: func() {
				mockSvc.EXPECT().Register(gomock.Any(), models.User{
					Username:    "username",
					Email:       "email@gmail.com",
					PhoneNumber: "phone_number",
					Fullname:    "fullname",
					Address:     "address",
					Dob:         "dob",
					Password:    "password",
				}).Return(models.User{}, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			r := gin.Default()

			h := &RegisterHandler{
				RegisterService: mockSvc,
			}

			userV1 := r.Group("/user/v1")
			userV1.POST("/register", h.Register)

			w := httptest.NewRecorder()

			endPoint := "/user/v1/register"
			model := models.User{
				Username:    "username",
				Email:       "email@gmail.com",
				PhoneNumber: "phone_number",
				Fullname:    "fullname",
				Address:     "address",
				Dob:         "dob",
				Password:    "password",
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
