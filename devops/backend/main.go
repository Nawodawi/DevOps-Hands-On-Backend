package main

import (
	"devops/backend/controllers"
	"devops/backend/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
