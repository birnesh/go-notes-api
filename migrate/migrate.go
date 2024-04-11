package main

import (
	"github.com/birnesh/notes/initializers"
	"github.com/birnesh/notes/models"
)

func init() {
	initializers.LoadEnvVarialbles()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Notes{})
	initializers.DB.AutoMigrate(&models.User{})
}
