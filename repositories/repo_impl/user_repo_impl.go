package repo_impl

import (
	"book/models"
	"book/repositories"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewInstanceUserRepo(db *gorm.DB) repositories.UserRepo {
	return &UserRepoImpl{db: db}
}

func (userRepo *UserRepoImpl) Save(user *models.User) error {
	return userRepo.db.Create(user).Error
}

func (userRepo *UserRepoImpl) GetByEmail(email string, user *models.User) error {
	return userRepo.db.Where("email = ?", email).First(user).Error
}

func (userRepo *UserRepoImpl) GetById(id string, user *models.User) error {
	return userRepo.db.Where("id = ?", id).First(user).Error
}
