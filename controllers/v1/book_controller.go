package controller_v1

import (
	"book/models"
	"book/repositories"
	"book/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct {
	BookRepo repositories.BookRepo
}


func (controller *BookController) UpdateFields(c *gin.Context){
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Params invalid"))
		return
	}

	var body map[string]interface{};
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return
	}


}

func (controller *BookController) UpdateAllField(c *gin.Context){
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Params invalid"))
		return
	}
	var book models.Book
	if err := c.ShouldBindBodyWithJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return
	}

	book.Id = id

	if err := controller.BookRepo.UpdateAllField(&book); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusBadRequest, models.MessageResponse("Data updated"))

}
func (controller *BookController) DeleteById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Params invalid"))
		return
	}
	
	if err := controller.BookRepo.DeleteById(id); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.MessageResponse("Delete record ("+id+") successful"))

}

func (controller *BookController) GetById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Params invalid"))
		return
	}
	var book models.Book
	if err := controller.BookRepo.GetById(id, &book); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.DataAndMessageResponse(book, "Get data successful"))

}

func (controller *BookController) GetAll(c *gin.Context) {
	var books []models.Book

	if err := controller.BookRepo.GetAll(&books); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.DataAndMessageResponse(books, "Get data successful"))

}

func (controller *BookController) Create(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindBodyWithJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return

	}
	if err := book.GenerateID(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(err.Error()))
		return
	}

	if err := utils.ValidateStruct(book); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return

	}
	if err := controller.BookRepo.Create(&book); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, models.DataResponse(book))
}
