package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tantm97/book-management-system-golang/pkg/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	authorized := incomingRoutes.Group("/api/v1")
	{
		authorized.POST("/users/signup/", controllers.Signup)
		authorized.POST("/users/login/", controllers.Login)
	}
}
