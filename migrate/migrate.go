package main

import (
	"github.com/birnesh/go-notes-api/initializers"
	"github.com/birnesh/go-notes-api/models"
)

func init() {
	initializers.LoadEnvVarialbles()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Notes{})
	initializers.DB.AutoMigrate(&models.User{})
}
