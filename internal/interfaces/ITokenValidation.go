package interfaces

import (
	"context"
	tokenValidation "ewallet-ums/cmd/proto/tokenvalidation"
	"ewallet-ums/helpers"
)

type ITokenValidationHandler interface {
	TokenValidationHandler(ctx context.Context, req *tokenValidation.TokenRequest) (*tokenValidation.TokenResponse, error)
	tokenValidation.UnimplementedTokenValidationServer
}

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}
