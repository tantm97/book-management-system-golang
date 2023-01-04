package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tantm97/book-management-system-golang/pkg/controllers"
	"github.com/tantm97/book-management-system-golang/pkg/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	authorized := incomingRoutes.Group("/api/v1")
	authorized.Use(middleware.Authenticate)
	{
		authorized.GET("/users/", controllers.GetUsers)
		authorized.PATCH("/users/:userId", controllers.UpdateUser)
	}
}
