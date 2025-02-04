package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id          string     `gorm:"size:255; primary_key; not null; unique" json:"id" validate:"uuid,required"`
	Name        string     `gorm:"not nul" json:"name" validate:"required"`
	Description string     `gorm:"not null" json:"description" validate:"required"`
	Author      string     `gorm:"not null" json:"author" validate:"required"`
	Year        int16     `gorm:"not null" json:"year" validate:"required,number,gte=500,lte=2050"`
	LastUpDate  time.Time `gorm:"autoUpdateTime" json:"last_up_date"`
}


func (book *Book) GenerateID() error{
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	book.Id = id.String()
	return nil
}