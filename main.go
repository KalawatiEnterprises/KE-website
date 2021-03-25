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
	var port string = "8080"
//	var port string = os.Getenv("PORT")
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*")

	// directories
	router.Static("/css", "static/css")
	router.Static("/scripts", "static/scripts")
	//router.Static("/fonts", "static/fonts/")
	//router.Static("/media", "static/media")

	// serve pages
	router.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/brands", func (c *gin.Context) {
		c.HTML(http.StatusOK, "brands.html", nil)
	})

	router.GET("/contact", func (c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})
	router.POST("/contact", getContact)

	router.Run(":" + port)
}

// this reads the context
// so nothing else needs to be passed
func getContact (c *gin.Context) {
	contact := contactInfo {
		name    : c.PostForm("fullname"),
		email   : c.PostForm("email"),
		phone   : c.PostForm("phone"),
		message : c.PostForm("message"),
	}

	/*
	if (contact.email != "") {
		confirmMail(contact.message, contact.email)
	}
	if (contact.name != "" && (contact.email != "" || contact.phone != "")) {
		notifySubmission(contact.name, contact.email, contact.phone, contact.message)
	}
	*/

	fmt.Println(contact) // for testing
	c.HTML(http.StatusOK, "thankyou.html", gin.H{
        	"name"    : contact.name,
        	"email"   : contact.email,
        	"phone"   : contact.phone,
		"message" : contact.message,
    	})
}
