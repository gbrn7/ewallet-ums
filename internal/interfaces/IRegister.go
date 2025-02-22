package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=IRegister.go -destination=../api/register_mock_test.go -package=api
type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}

type IRegisterHandler interface {
	Register(*gin.Context)
}
