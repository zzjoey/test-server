package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200,
			"Hello Test Server!",
		)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
