package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=ILogout.go -destination=../api/logout_mock_test.go -package=api
type ILogoutService interface {
	Logout(ctx context.Context, token string) error
}

type ILogoutHandler interface {
	Logout(*gin.Context)
}
