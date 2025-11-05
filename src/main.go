package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong", "request_id": uuid.New().String()})
	})
	router.Run(":9000")

}
