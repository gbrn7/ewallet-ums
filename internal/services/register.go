package services

import (
	"context"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepo       interfaces.IUserRepository
	ExternalWallet interfaces.IWallet
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashedPassword)

	err = s.UserRepo.InsertNewUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	_, err = s.ExternalWallet.CreateWallet(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	resp := request
	resp.Password = ""
	return resp, nil
}
