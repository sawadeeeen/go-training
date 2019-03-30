package main

import (
	"github.com/gin-gonic/gin"
)

// 同じディレクトリ内には、mainは一つしか持てない
func aain() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")

	})
	r.Run()
}
