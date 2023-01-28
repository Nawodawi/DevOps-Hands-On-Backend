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

func PostsShow(c *gin.Context) {
	//Get id of url
	id := c.Param("id")
	//Get posts
	var post models.Post
	initializers.DB.First(&post, id)
	//Response with them

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	//Get the id from the url
	id := c.Param("id")
	//Get the data from the req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//Get the id of the url
	id := c.Param("id")
	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	//Response
	c.Status(200)

}
