package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Serve static files
	r.Static("/assets", "../../assets")

	// Load templates
	r.LoadHTMLGlob("../../templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "My Portfolio",
		})
	})

	return r
}
