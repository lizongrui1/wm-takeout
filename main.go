package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Auth endpoint reached",
		})
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
