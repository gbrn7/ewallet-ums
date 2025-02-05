package interfaces

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type IRefrehTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error)
}

type IRefreshTokenHandler interface {
	RefreshToken(*gin.Context)
}
