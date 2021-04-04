package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"os"
	"fmt"
)

type contactInfo struct {
	name string
	email string
	phone string
	message string
}

func main() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)
	var port string = "8080"
//	var port string = os.Getenv("PORT")
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*")

	// directories
	router.Static("/css", "static/css")
	router.Static("/scripts", "static/scripts")
	router.Static("/media", "static/media")

	router.Static("/catalogues", "static/media/catalogues")

	// serve pages
	router.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/brands", func (c *gin.Context) {
		c.HTML(http.StatusOK, "brands.html", nil)
	})

	router.GET("/comingsoon", func (c *gin.Context) {
		c.HTML(http.StatusOK, "coming-soon.html", nil)
	})

	router.GET("/catalogues", func (c *gin.Context) {
		c.String(http.StatusOK, "coming soon")
	})

	router.POST("/contact", getContact)

	router.Run(":" + port)
}

func getContact (c *gin.Context) {
	contact := contactInfo {
		name    : c.PostForm("fullname"),
		email   : c.PostForm("email"),
		phone   : c.PostForm("phone"),
		message : c.PostForm("message"),
	}

	/*
	if (contact.name != "" && (contact.email != "" || contact.phone != "")) {
		responseMail(contact.name, contact.message, contact.email)
		//notifySubmission(contact.name, contact.email, contact.phone, contact.message)
	}
	*/

	fmt.Println(contact) // for testing
	/*
	c.HTML(200, "thankyou.html", gin.H {
         	"name"    : contact.name,
         	"email"   : contact.email,
         	"phone"   : contact.phone,
	 	"message" : contact.message,
    	})
	*/
	// send confirmation to the client
	c.JSON(200, "ye")
}
