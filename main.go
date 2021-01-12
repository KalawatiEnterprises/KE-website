package main

import (
	"github.com/gin-gonic/gin"
	//"log"
	"net/http"
	//"os"
)

var port string
func getPort() { // automatically get the assigned port on heroku
	port = "8069"
//	port = os.Getenv("PORT")

//	if port == "" {
//		log.Fatal("$PORT must be set")
//	}
}

func main() {
	getPort()
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*")

	// directories
	router.Static("/static", "static")
	router.Static("/css", "static/css")
	router.Static("/scripts", "static/scripts")
	router.Static("/media", "static/media")

	// serve pages
	router.GET("/", func (c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/brands", func (c *gin.Context) {
	c.HTML(http.StatusOK, "brands.html", nil)
	})

	router.GET("/contactus", func (c *gin.Context) {
	c.HTML(http.StatusOK, "contactus.html", nil)
	})

	router.Run(":" + port)
}
