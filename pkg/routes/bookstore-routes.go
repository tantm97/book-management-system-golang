package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tantm97/book-management-system-golang/pkg/controllers"
	"github.com/tantm97/book-management-system-golang/pkg/middleware"
)

func BookStoreRoutes(incomingRoutes *gin.Engine) {
	authorized := incomingRoutes.Group("/api/v1")
	authorized.Use(middleware.Authenticate)
	{
		authorized.POST("/book/", controllers.CreateBook)
		authorized.GET("/book/", controllers.GetBooks)
		authorized.GET("/book/:bookId", controllers.GetBookById)
		authorized.PATCH("/book/:bookId", controllers.UpdateBook)
		authorized.DELETE("/book/:bookId", controllers.DeleteBook)
	}
}
