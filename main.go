package main

import "github.com/gin-gonic/gin"

func wolfHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "software development in go",
	})
}

func main() {
	router := gin.Default()
	router.GET("/wolf", wolfHandler)
	router.Run(":3010")
}
