package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/wolf", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "software development in go",
		})
	})
	router.Run(":3010")
}
