package main

import (
	"github.com/gin-gonic/gin"

	"github.com/DigdarshanMohanty/notes-api/internal/db"
	"github.com/DigdarshanMohanty/notes-api/internal/handlers/notes"
	"github.com/DigdarshanMohanty/notes-api/internal/handlers/users"
	"github.com/DigdarshanMohanty/notes-api/internal/middleware"
	"github.com/DigdarshanMohanty/notes-api/internal/models"
)

func main() {
	d := db.Connect()
	models.Migrate(d)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", d)
		c.Next()
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/register", users.Register)
	r.POST("/login", users.Login)

	auth := r.Group("/")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/user/:id", users.ShowUserDetails)
		auth.POST("/notes", notes.CreateNote)
		auth.GET("/notes", notes.GetNotes)
		auth.DELETE("/notes/:id", notes.DeleteNote)
		auth.PATCH("/notes/:id", notes.UpdateNote)
	}

	r.Run(":8080") 
}
