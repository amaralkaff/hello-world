package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//initial Gin
	router := gin.Default()

	// Routes method GET
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Server port 3000
	router.Run(":3000")
}
