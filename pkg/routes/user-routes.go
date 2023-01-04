package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tantm97/book-management-system-golang/pkg/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users/", controllers.GetUsers)
	incomingRoutes.PATCH("/users/:userId", controllers.UpdateUser)
}
