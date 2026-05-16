package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("web/templates/*")

	// Static files
	r.Static("/static", "./web/static")

	// Route
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "kobraa.chess",
		})
	})

	// Run server
	r.Run(":8080")
}