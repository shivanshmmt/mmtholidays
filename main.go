package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shivanshmmt/holidays/controllers"
	//"net/http"
	"github.com/shivanshmmt/holidays/models"
	//"github.com/rahmanfadhil/gin-bookstore/controllers/books.go"
)

func main() {
	r := gin.Default()

	//r.GET("/check", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	//})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)

	models.ConnectDatabase()


	r.Run()
}