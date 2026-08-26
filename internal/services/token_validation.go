package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"

	"github.com/pkg/errors"
)

type TokenValidationService struct {
	UserRepo interfaces.IUserRepository
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error) {

	var (
		claimToken *helpers.ClaimToken
		err        error
	)

	claimToken, err = helpers.ValidateToken(ctx, token)
	if err != nil {
		helpers.Logger.Warn("Validate Token Failed")
		return claimToken, errors.Wrap(err, "failed to validate token")
	}

	_, err = s.UserRepo.GetUserSessionByToken(ctx, token)
	if err != nil {
		helpers.Logger.Warn("failed to get user session")
		return claimToken, errors.Wrap(err, "failed to get user session")
	}

	helpers.Logger.Info("Validate Token Succes")

	return claimToken, nil
}
