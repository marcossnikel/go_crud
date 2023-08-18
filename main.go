package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcossnikel/go_api/controllers"
	"github.com/marcossnikel/go_api/models"
)

func main() {
	router := gin.Default() // router with default middleware installed

	models.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPostById)
	router.PATCH("/posts/:id", controllers.UpdatePostById)
	router.DELETE("/posts/:id", controllers.DeletePostById)

	router.Run()
}
