package interfaces

import (
	"context"
	"ewallet-ums/internal/models"
)

//go:generate mockgen -source=IUser.go -destination=../services/user_mock_test.go -package=services
type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	DeleteUserSession(ctx context.Context, token string) error
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error)
	UpdateTokenByRefreshToken(ctx context.Context, token string, refreshToken string) error
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error)
}
