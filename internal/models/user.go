package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username" gorm:"column:username;type:varchar(20)"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(100)"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(15)"`
	Fullname    string    `json:"full_name" gorm:"column:full_name;type:varchar(100)"`
	Address     string    `json:"address" gorm:"column:address;type:text"`
	Dob         string    `json:"dob" gorm:"column:dob;type:date"`
	Password    string    `json:"password,omitempty" gorm:"column:password;type:varchar(255)"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (*User) TableName() string {
	return "Users"
}
func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  uint `gorm:"primarykey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              int       `json:"user_id" gorm:"type:int" validate:"required"`
	Token               string    `json:"token" gorm:"type:varchar(255)" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:varchar(255)" validate:"required"`
	TokenExpired        time.Time `json:"token_expired" validate:"required"`
	RefreshTokenExpired time.Time `json:"refresh_token_expired" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
