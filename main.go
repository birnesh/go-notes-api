package main

import (
	"github.com/birnesh/go-notes-api/controllers"
	"github.com/birnesh/go-notes-api/initializers"
	"github.com/birnesh/go-notes-api/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVarialbles()
	initializers.ConnectToDb()
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/note", middleware.RequireAuth, controllers.PostNote)
	r.GET("/note", middleware.RequireAuth, controllers.ListNote)
	r.GET("/note/:id", middleware.RequireAuth, controllers.FetchNote)
	r.PUT("/note/:id", middleware.RequireAuth, controllers.UpdateNote)
	r.DELETE("/note/:id", middleware.RequireAuth, controllers.DeleteNote)
	r.Run() // listen and serve on 0.0.0.0:8080
}
