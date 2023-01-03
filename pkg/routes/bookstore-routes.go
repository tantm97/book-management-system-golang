package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tantm97/book-management-system-golang/pkg/controllers"
)

func BookStoreRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/book/", controllers.CreateBook)
	incomingRoutes.GET("/book/", controllers.GetBooks)
	incomingRoutes.GET("/book/:bookId", controllers.GetBookById)
	incomingRoutes.PATCH("/book/:bookId", controllers.UpdateBook)
	incomingRoutes.DELETE("/book/:bookId", controllers.DeleteBook)
}
