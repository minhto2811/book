package models

import (
	"book/security"
	"fmt"

	"github.com/google/uuid"
)

const (
	USER int16 = iota
	ADMIN
)

type User struct {
	Id   string `gorm:"primary_key; not null; unique" json:"id"`
	Name string `gorm:"size:255; not null" json:"name"`
	Role int16  `gorm:"not null" json:"role"`
	Account
}

type Account struct {
	Email    string `gorm:"size:255; not null; unique" json:"email"`
	Password string `gorm:"size:255; not null" json:"password"`
}

type UserInfo struct {
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

func (user *User) GenarateID() error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	user.Id = id.String()
	return nil
}

func (user *User) HashPassword() error {
	result, err := security.GenerateFromPassword(user.Password)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	user.Password = string(result)
	return nil

}

func (user *User) ComparePassword(password string) bool {
	return security.ComparePassword(user.Password, password)
}
