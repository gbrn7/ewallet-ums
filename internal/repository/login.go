package repository

import (
	"context"
	"ewallet-ums/internal/models"

	"gorm.io/gorm"
)

type LoginRepository struct {
	DB *gorm.DB
}

func (r *LoginRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var (
		user models.User
		err  error
	)

	err = r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, nil
	}
	return user, nil
}
