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

func TestUserRepository_InsertNewUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	assert.NoError(t, err)

	type args struct {
		ctx   context.Context
		model models.User
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
				model: models.User{
					Email:       "test@gmail.com",
					Username:    "raygbrn",
					PhoneNumber: "12345678",
					Fullname:    "Muhammad Rayhan Gibran",
					Address:     "Malang",
					Dob:         "16/03/2003",
					Password:    "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`username`,`email`,`phone_number`,`full_name`,`address`,`dob`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?)")).WithArgs(
					args.model.Username,
					args.model.Email,
					args.model.PhoneNumber,
					args.model.Fullname,
					args.model.Address,
					args.model.Dob,
					args.model.Password,
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				model: models.User{
					Email:       "test@gmail.com",
					Username:    "raygbrn",
					PhoneNumber: "12345678",
					Fullname:    "Muhammad Rayhan Gibran",
					Address:     "Malang",
					Dob:         "16/03/2003",
					Password:    "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`username`,`email`,`phone_number`,`full_name`,`address`,`dob`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?)")).WithArgs(
					args.model.Username,
					args.model.Email,
					args.model.PhoneNumber,
					args.model.Fullname,
					args.model.Address,
					args.model.Dob,
					args.model.Password,
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &UserRepository{
				DB: gormDB,
			}
			if err := r.InsertNewUser(tt.args.ctx, &tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.InsertNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
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

			r := &UserRepository{
				DB: gormDB,
			}
			got, err := r.GetUserByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserByUsername() = %v, want %v", got, tt.want)
			}

			assert.NoError(t, mock.ExpectationsWereMet())

		})
	}
}

func TestUserRepository_InsertNewUserSession(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	assert.NoError(t, err)

	now := time.Now()
	tokenExpired := now.AddDate(0, 0, 7)
	refreshTokenExpired := now.AddDate(0, 0, 28)

	type args struct {
		ctx   context.Context
		model *models.UserSession
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
				ctx: context.Background(),
				model: &models.UserSession{
					UserID:              1,
					Token:               "token",
					RefreshToken:        "refresh_token",
					TokenExpired:        tokenExpired,
					RefreshTokenExpired: refreshTokenExpired,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `user_sessions` (`created_at`,`updated_at`,`user_id`,`token`,`refresh_token`,`token_expired`,`refresh_token_expired`) VALUES (?,?,?,?,?,?,?)")).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.UserID,
					args.model.Token,
					args.model.RefreshToken,
					args.model.TokenExpired,
					args.model.RefreshTokenExpired,
				).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				model: &models.UserSession{
					UserID:              1,
					Token:               "token",
					RefreshToken:        "refresh_token",
					TokenExpired:        time.Now().AddDate(0, 0, 7),
					RefreshTokenExpired: time.Now().AddDate(0, 0, 28),
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `user_sessions` (`created_at`,`updated_at`,`user_id`,`token`,`refresh_token`,`token_expired`,`refresh_token_expired`) VALUES (?,?,?,?,?,?,?)")).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.UserID,
					args.model.Token,
					args.model.RefreshToken,
					args.model.TokenExpired,
					args.model.RefreshTokenExpired,
				).WillReturnError(assert.AnError)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &UserRepository{
				DB: gormDB,
			}
			if err := r.InsertNewUserSession(tt.args.ctx, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.InsertNewUserSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		assert.NoError(t, mock.ExpectationsWereMet())
	}
}

func TestUserRepository_DeleteUserSession(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	assert.NoError(t, err)

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
				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `user_sessions` WHERE token = ?")).WithArgs(
					args.token,
				).WillReturnResult(sqlmock.NewResult(1, 1))
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
				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `user_sessions` WHERE token = ?")).WithArgs(
					args.token,
				).WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &UserRepository{
				DB: gormDB,
			}
			if err := r.DeleteUserSession(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.DeleteUserSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		assert.NoError(t, mock.ExpectationsWereMet())

	}
}

func TestUserRepository_UpdateTokenByRefreshToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	assert.NoError(t, err)

	type args struct {
		ctx          context.Context
		token        string
		refreshToken string
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
				ctx:          context.Background(),
				token:        "token",
				refreshToken: "refresh_token",
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `user_sessions` SET token = ? WHERE refresh_token = ?")).WithArgs(
					args.token,
					args.refreshToken,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "error",
			args: args{
				ctx:          context.Background(),
				token:        "token",
				refreshToken: "refresh_token",
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `user_sessions` SET token = ? WHERE refresh_token = ?")).WithArgs(
					args.token,
					args.refreshToken,
				).WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				DB: gormDB,
			}
			if err := r.UpdateTokenByRefreshToken(tt.args.ctx, tt.args.token, tt.args.refreshToken); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.UpdateTokenByRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		assert.NoError(t, mock.ExpectationsWereMet())
	}
}

func TestUserRepository_GetUserSessionByToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	assert.NoError(t, err)

	now := time.Now()
	tokenExpired := now.AddDate(0, 0, 7)
	refreshTokenExpired := now.AddDate(0, 0, 28)

	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    models.UserSession
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx:   context.Background(),
				token: "token",
			},
			want: models.UserSession{
				ID:                  1,
				CreatedAt:           now,
				UpdatedAt:           now,
				UserID:              1,
				Token:               "token",
				RefreshToken:        "refresh_token",
				TokenExpired:        tokenExpired,
				RefreshTokenExpired: refreshTokenExpired,
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_sessions` WHERE token = ? ORDER BY `user_sessions`.`id` LIMIT ?")).WithArgs(
					args.token,
					1,
				).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "user_id", "token", "refresh_token", "token_expired", "refresh_token_expired"}).AddRow(1, now, now, 1, "token", "refresh_token", tokenExpired, refreshTokenExpired))
			},
		},
		{
			name: "error",
			args: args{
				ctx:   context.Background(),
				token: "token",
			},
			want:    models.UserSession{},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_sessions` WHERE token = ? ORDER BY `user_sessions`.`id` LIMIT ?")).WithArgs(
					args.token,
					1,
				).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				DB: gormDB,
			}
			got, err := r.GetUserSessionByToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserSessionByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserSessionByToken() = %v, want %v", got, tt.want)
			}
		})

		assert.NoError(t, mock.ExpectationsWereMet())
	}
}

func TestUserRepository_GetUserSessionByRefreshToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	assert.NoError(t, err)

	now := time.Now()
	tokenExpired := now.AddDate(0, 0, 7)
	refreshTokenExpired := now.AddDate(0, 0, 28)

	type args struct {
		ctx          context.Context
		refreshToken string
	}
	tests := []struct {
		name    string
		args    args
		want    models.UserSession
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx:          context.Background(),
				refreshToken: "refresh_token",
			},
			want: models.UserSession{
				ID:                  1,
				CreatedAt:           now,
				UpdatedAt:           now,
				UserID:              1,
				Token:               "token",
				RefreshToken:        "refresh_token",
				TokenExpired:        tokenExpired,
				RefreshTokenExpired: refreshTokenExpired,
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_sessions` WHERE refresh_token = ? ORDER BY `user_sessions`.`id` LIMIT ?")).WithArgs(
					args.refreshToken,
					1,
				).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "user_id", "token", "refresh_token", "token_expired", "refresh_token_expired"}).AddRow(1, now, now, 1, "token", "refresh_token", tokenExpired, refreshTokenExpired))
			},
		},
		{
			name: "error",
			args: args{
				ctx:          context.Background(),
				refreshToken: "refresh_token",
			},
			want:    models.UserSession{},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_sessions` WHERE refresh_token = ? ORDER BY `user_sessions`.`id` LIMIT ?")).WithArgs(
					args.refreshToken,
					1,
				).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &UserRepository{
				DB: gormDB,
			}
			got, err := r.GetUserSessionByRefreshToken(tt.args.ctx, tt.args.refreshToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserSessionByRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserSessionByRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
