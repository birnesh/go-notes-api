package controllers

import (
	"net/http"

	"github.com/birnesh/notes/initializers"
	"github.com/birnesh/notes/models"
	"github.com/gin-gonic/gin"
)

func PostNote(c *gin.Context) {
	// get the notes body

	var body struct {
		Title       string
		IsCompleted bool
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}
	// crete notes
	note := models.Notes{Title: body.Title, IsCompleted: body.IsCompleted}

	result := initializers.DB.Create(&note)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{})
}

func ListNote(c *gin.Context) {
	// get all notes
	var notes []models.Notes
	initializers.DB.Find(&notes)

	// response
	c.JSON(http.StatusOK, gin.H{"data": notes})
}
func FetchNote(c *gin.Context) {

	// get id
	id := c.Param("id")
	// get all notes
	var note models.Notes
	initializers.DB.Find(&note, id)

	// response
	c.JSON(http.StatusOK, gin.H{"data": note})
}
func UpdateNote(c *gin.Context) {

	// get id
	id := c.Param("id")

	var body struct {
		Title       string
		IsCompleted bool
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	// get the note
	var note models.Notes
	initializers.DB.Find(&note, id)

	// update the note
	note.Title = body.Title
	note.IsCompleted = body.IsCompleted
	initializers.DB.Save(&note)

	// response
	c.JSON(http.StatusOK, gin.H{"data": note})
}
func DeleteNote(c *gin.Context) {

	// get id
	id := c.Param("id")

	var body struct {
		Title       string
		IsCompleted bool
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	// delete note
	initializers.DB.Delete(&models.Notes{}, id)

	// response
	c.JSON(http.StatusOK, gin.H{})
}
