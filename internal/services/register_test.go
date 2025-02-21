package services

import (
	"context"
	"ewallet-ums/internal/models"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterService_Register(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	userMockRepo := NewMockIUserRepository(ctrlMock)
	externalMockRepo := NewMockIExternal(ctrlMock)

	type args struct {
		ctx     context.Context
		request models.User
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				request: models.User{
					Username:    "username",
					Password:    "password",
					Fullname:    "username",
					Email:       "username@gmail.com",
					PhoneNumber: "1212122",
					Address:     "indonesia",
					Dob:         "1990-10-10",
				},
			},
			want: models.User{
				Username:    "username",
				Fullname:    "username",
				Email:       "username@gmail.com",
				PhoneNumber: "1212122",
				Address:     "indonesia",
				Dob:         "1990-10-10",
			},
			wantErr: false,
			mockFn: func(args args) {
				userMockRepo.EXPECT().InsertNewUser(args.ctx, gomock.Any()).Return(nil)
				externalMockRepo.EXPECT().CreateWallet(args.ctx, gomock.Any()).Return(nil, nil)

				notifMsg := map[string]string{
					"full_name": args.request.Fullname,
				}

				externalMockRepo.EXPECT().SendNotification(args.ctx, args.request.Email, "register", notifMsg).Return(nil)
			},
		},
		{
			name: "failed when insert new user",
			args: args{
				ctx: context.Background(),
				request: models.User{
					Username:    "username",
					Password:    "password",
					Fullname:    "username",
					Email:       "username@gmail.com",
					PhoneNumber: "1212122",
					Address:     "indonesia",
					Dob:         "1990-10-10",
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				userMockRepo.EXPECT().InsertNewUser(args.ctx, gomock.Any()).Return(assert.AnError)
			},
		},
		{
			name: "failed when send notification",
			args: args{
				ctx: context.Background(),
				request: models.User{
					Username:    "username",
					Password:    "password",
					Fullname:    "username",
					Email:       "username@gmail.com",
					PhoneNumber: "1212122",
					Address:     "indonesia",
					Dob:         "1990-10-10",
				},
			},
			want: models.User{
				Username:    "username",
				Fullname:    "username",
				Email:       "username@gmail.com",
				PhoneNumber: "1212122",
				Address:     "indonesia",
				Dob:         "1990-10-10",
			},
			wantErr: false,
			mockFn: func(args args) {
				userMockRepo.EXPECT().InsertNewUser(args.ctx, gomock.Any()).Return(nil)
				externalMockRepo.EXPECT().CreateWallet(args.ctx, gomock.Any()).Return(nil, nil)

				notifMsg := map[string]string{
					"full_name": args.request.Fullname,
				}

				externalMockRepo.EXPECT().SendNotification(args.ctx, args.request.Email, "register", notifMsg).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &RegisterService{
				UserRepo: userMockRepo,
				External: externalMockRepo,
			}
			got, err := s.Register(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterService.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
