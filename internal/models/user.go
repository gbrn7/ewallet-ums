package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          uint64    `json:"id"`
	Username    string    `json:"username" gorm:"column:username;type:varchar(20);unique;not null" validate:"required"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(100);unique;not null" validate:"required"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(15);unique;not null" validate:"required"`
	Fullname    string    `json:"full_name" gorm:"column:full_name;type:varchar(100)" validate:"required"`
	Address     string    `json:"address" gorm:"column:address;type:text"`
	Dob         string    `json:"dob" gorm:"column:dob;type:date"`
	Password    string    `json:"password,omitempty" gorm:"column:password;type:varchar(255)" validate:"required"`
	CreatedAt   time.Time `json:"-"  gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time `json:"-"  gorm:"column:updated_at;type:timestamp"`
}

func (*User) TableName() string {
	return "users"
}
func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  uint `gorm:"primarykey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              uint64    `json:"user_id" gorm:"type:int" validate:"required"`
	Token               string    `json:"token" gorm:"type:text" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:text" validate:"required"`
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
