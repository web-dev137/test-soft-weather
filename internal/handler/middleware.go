package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	if c.Request.Header.Get("User-Access") == "superuser" {
		c.Next()
	} else {
		fmt.Print("user unauthorized")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user unauthorized"})
		c.Abort()
	}
}
