package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tantm97/book-management-system-golang/pkg/database"
	"github.com/tantm97/book-management-system-golang/pkg/routes"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	database.Connect()

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.BookStoreRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
