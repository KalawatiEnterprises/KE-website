package main

import (
	"github.com/gin-gonic/gin"
	//"log"
	"net/http"
	//"os"
	//"fmt"
)

var port string

func getPort() { // automatically get the assigned port on heroku
	port = "8080"
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
	router.Static("/static", "static")

	// serve pages
	router.GET("/", homePage)
	router.GET("/locate", locatePage)


	router.Run(":" + port)
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func locatePage(c *gin.Context) {
        c.HTML(http.StatusOK, "locate-us.html", nil)
}

