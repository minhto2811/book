package repo_impl

import (
	"book/models"
	"book/repositories"

	"gorm.io/gorm"
)

type BookRepoImpl struct {
	db *gorm.DB
}



func NewInstanceBookRepo(db *gorm.DB) repositories.BookRepo {
	return &BookRepoImpl{db: db}
}

func (bookRepo *BookRepoImpl)UpdateAllField(book *models.Book) error{
	return bookRepo.db.Updates(book).Error
}

func (bookRepo *BookRepoImpl) Create(book *models.Book) error {
	return bookRepo.db.Create(book).Error
}

func (bookRepo *BookRepoImpl) GetAll(books *[]models.Book) error {
	return bookRepo.db.Find(books).Error
}

func (bookRepo *BookRepoImpl) GetById(id string, book *models.Book) error {
	return bookRepo.db.First(book, map[string]interface{}{"id": id}).Error
}

func (bookRepo *BookRepoImpl) DeleteById(id string) error {
	return bookRepo.db.Delete(&models.Book{}, map[string]interface{}{"id": id}).Error
}
