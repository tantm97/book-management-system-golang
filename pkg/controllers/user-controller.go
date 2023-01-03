package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tantm97/book-management-system-golang/pkg/database"
	"github.com/tantm97/book-management-system-golang/pkg/models"
)

func HashPassword(password string) string {

}

func Signup(c *gin.Context) {
	var input models.UserSignup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count uint64
	result := database.DB.Model(&models.User{}).Where("email=?", input.Email).Count(&count)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email is existed."})
		return
	}

	result = database.DB.Model(&models.User{}).Where("email=?", input.Phone).Count(&count)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This phone number is existed."})
		return
	}

	newUser := models.User{
		FirstName: input.FirstName,
		LastName: input.LastName,
		Phone: input.Phone,
	}
	newUser.Password = HashPassword(input.Password)
	token, refreshToken, _ := 



}

func Login(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}
