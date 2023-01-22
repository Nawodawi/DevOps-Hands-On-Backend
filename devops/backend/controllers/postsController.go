package controllers

import (
	"devops/backend/initializers"
	"devops/backend/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body

	//Create a post
	post := models.Post{Title: "Hello", Body: "this is a body"}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}
