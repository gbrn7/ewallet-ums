package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=ILogin.go -destination=../api/login_mock_test.go -package=api
type ILoginService interface {
	Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error)
}

type ILoginHandler interface {
	Login(c *gin.Context)
}
