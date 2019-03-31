package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.Use(sampleMiddleware)
	route.GET("/", func(c *gin.Context) {
		log.Println("main logic")
		c.JSON(200, gin.H{"message": "Hello"})
	})
	route.Run(":8070")
}

var sampleMiddleware = func(c *gin.Context) {
	log.Println("before logic")
	c.Next()
	log.Println("after logic")
}

// func sampleMiddleware (c *gin.Context) {
// 	log.Println("before logic")
// 	c.Next()
// 	log.Println("after logic")
// }

// func sampleMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		log.Println("before logic")
// 		c.Next()
// 		log.Println("after logic")
// 	}
// }
