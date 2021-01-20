package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"os"
	"fmt"
)

var (
	port string
)

func getPort() { // automatically get the assigned port on heroku
	port = "8080"
//	port = os.Getenv("PORT")

//	if port == "" {
//		log.Fatal("$PORT must be set")
//	}
}

type contactInfo struct {
	name string
	email string
	phone string
	message string
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

	//router.POST("/contactus", getContact)

	router.Run(":" + port)
}

func getContact (c *gin.Context) {
	return
	contact := contactInfo {
		name    : c.PostForm("fullname"),
		email   : c.PostForm("email"),
		phone   : c.PostForm("phone"),
		message : c.PostForm("message"),
	}
	if (contact.email != "") {
		confirmMail(contact.message, contact.email)
	}
	if (contact.name != "" && (contact.email != "" || contact.phone != "")) {
		notifySubmission(contact.name, contact.email, contact.phone, contact.message)
	}
	fmt.Println(contact)
}
