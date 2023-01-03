package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tantm97/book-management-system-golang/pkg/database"
	"github.com/tantm97/book-management-system-golang/pkg/models"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	result := database.DB.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	}
	c.JSON(http.StatusOK, books)
}

func GetBookById(c *gin.Context) {
	var book models.Book
	bookId := c.Param("bookId")
	if err := database.DB.Where("id=?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found."})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var input models.CreateBook
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{Title: input.Title,
		Author:      input.Author,
		Publication: input.Publication}
	result := database.DB.Create(&book)
	if result.Error != nil {
		msg := fmt.Sprint("Book was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	bookId := c.Param("bookId")
	if err := database.DB.Where("id=?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found."})
		return
	}
	result := database.DB.Delete(&book)
	if result.Error != nil {
		msg := fmt.Sprint("Book was not deleted")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdateBook(c *gin.Context) {
	var input models.UpdateBook
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var foundedBook models.Book
	bookId := c.Param("bookId")
	if err := database.DB.Where("id=?", bookId).First(&foundedBook).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found."})
		return
	}

	if input.Author != "" {
		foundedBook.Author = input.Author
	}
	if input.Publication != "" {
		foundedBook.Publication = input.Publication
	}
	if input.Title != "" {
		foundedBook.Title = input.Title
	}
	result := database.DB.Save(&foundedBook)
	if result.Error != nil {
		msg := fmt.Sprint("Book was not updated")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, foundedBook)
}
