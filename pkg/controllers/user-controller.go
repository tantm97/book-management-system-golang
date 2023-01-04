package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tantm97/book-management-system-golang/pkg/database"
	"github.com/tantm97/book-management-system-golang/pkg/helpers"
	"github.com/tantm97/book-management-system-golang/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = fmt.Sprint("Password is incorrect.")
		check = false
	}
	return check, msg
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

	result = database.DB.Model(&models.User{}).Where("phone=?", input.Phone).Count(&count)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This phone number is existed."})
		return
	}

	newUser := models.User{
		FirstName: &input.FirstName,
		LastName:  &input.LastName,
		Phone:     &input.Phone,
		Email:     &input.Email,
	}
	password := HashPassword(input.Password)
	newUser.Password = &password

	result = database.DB.Create(&newUser)
	if result.Error != nil {
		msg := fmt.Sprint("user was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	if err := database.DB.Where("email=?", input.Email).First(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, refreshToken, _ := helpers.GenerateAllTokens(*newUser.Email, *newUser.FirstName, *newUser.LastName, strconv.FormatInt(int64(newUser.ID), 10))
	newUser.Token = &token
	newUser.RefreshToken = &refreshToken

	result = database.DB.Save(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func Login(c *gin.Context) {
	var input models.UserLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var foundedUser models.User
	if err := database.DB.Where("email=?", input.Email).First(&foundedUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found."})
		return
	}
	passwordIsValid, msg := VerifyPassword(input.Password, *foundedUser.Password)
	if passwordIsValid != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := helpers.GenerateAllTokens(*foundedUser.Email, *foundedUser.FirstName, *foundedUser.LastName, strconv.FormatInt(int64(foundedUser.ID), 10))
	foundedUser.Token = &token
	foundedUser.RefreshToken = &refreshToken

	result := database.DB.Save(&foundedUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, foundedUser)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	}
	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	var input models.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var foundedUser models.User
	userId := c.Param("userId")

	if err := database.DB.Where("id=?", userId).First(&foundedUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
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

	result = database.DB.Model(&models.User{}).Where("phone=?", input.Phone).Count(&count)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This phone number is existed."})
		return
	}

	if input.Email != "" {
		foundedUser.Email = &input.Email
	}
	if input.FirstName != "" {
		foundedUser.FirstName = &input.FirstName
	}
	if input.LastName != "" {
		foundedUser.LastName = &input.LastName
	}
	if input.Phone != "" {
		foundedUser.Phone = &input.Phone
	}
	if input.Password != "" {
		password := HashPassword(input.Password)
		foundedUser.Phone = &password
	}

	result = database.DB.Save(&foundedUser)
	if result.Error != nil {
		msg := fmt.Sprint("User was not updated")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, foundedUser)
}
