package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tantm97/book-management-system-golang/pkg/helpers"
)

func Authenticate(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprint("No authorization header provides.")})
		c.Abort()
		return
	}
	claims, err := helpers.ValidateToken(clientToken)
	if err != "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		c.Abort()
		return
	}
	c.Set("email", claims.Email)
	c.Set("first_name", claims.FirstName)
	c.Set("last_name", claims.LastName)
	c.Set("uid", claims.Uid)
	c.Next()
}
