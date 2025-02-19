package repository

import (
	"context"
	"ewallet-ums/internal/models"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestLoginRepository_GetUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	assert.NoError(t, err)

	now := time.Now()

	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				username: "raygbrn",
			},
			want: models.User{
				ID:          1,
				Username:    "raygbrn",
				Email:       "test@gmail.com",
				PhoneNumber: "082132345432",
				Address:     "Malang",
				Fullname:    "gibran",
				Dob:         "16/03/2003",
				Password:    "password",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT ?")).WithArgs(
					args.username,
					1,
				).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "phone_number", "full_name", "address", "dob", "password", "created_at", "updated_at"}).AddRow(1, "raygbrn", "test@gmail.com", "082132345432", "gibran", "Malang", "16/03/2003", "password", now, now))
			},
		},
		{
			name: "error",
			args: args{
				ctx:      context.Background(),
				username: "raygbrn",
			},
			want:    models.User{},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT ?")).WithArgs(
					args.username,
					1,
				).WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &LoginRepository{
				DB: gormDB,
			}
			got, err := r.GetUserByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginRepository.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginRepository.GetUserByUsername() = %v, want %v", got, tt.want)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
