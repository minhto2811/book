package repositories

import "book/models"


type UserRepo interface {
	Save(user *models.User) error
	GetByEmail(email string,user *models.User)  error
	GetById(id string, user *models.User) error
}