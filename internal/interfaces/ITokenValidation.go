package interfaces

import (
	"context"
	"ewallet-ums/cmd/proto/tokenvalidation"
	"ewallet-ums/helpers"
)

//go:generate mockgen -source=ITokenValidation.go -destination=../api/token_mock_test.go -package=api
type ITokenValidationHandler interface {
	ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error)
}

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}
