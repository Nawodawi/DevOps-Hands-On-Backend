package controllers

import (
	"devops/backend/initializers"
	"devops/backend/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

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

func PostsIndex(c *gin.Context) {
	//Get posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//Response with them

	c.JSON(200, gin.H{
		"posts": posts,
	})

}
