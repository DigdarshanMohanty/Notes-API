package handlers

import (
	"net/http"

	"github.com/DigdarshanMohanty/notes-api/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.GetUint("user_id")

	var body struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	note := models.Note{
		Title:   body.Title,
		Content: body.Content,
		UserID:  userID,
	}
	if err := db.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create note"})
		return
	}
	db.Preload("User").First(&note, note.ID)
	c.JSON(http.StatusCreated, note)
}

func GetNotes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.GetUint("user_id")

	var notes []models.Note
	if err := db.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch notes"})
		return
	}

	if len(notes) == 0 {
		c.Status(http.StatusNoContent)
		return
	}
	c.JSON(http.StatusOK, notes)
}

func DeleteNote(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	userID := c.GetUint("user_id")
	noteID := c.Param("id")

	var note models.Note
	if err := db.Where("id = ? AND  user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}
	if err := db.Delete(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete note"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted successfully",
	})
}

func UpdateNote(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	userID := c.GetUint("user_id")
	noteID := c.Param("id")

	var body struct {
		Title string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var note models.Note
	if err := db.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	if body.Title != "" {
		note.Title = body.Title
	}
	if body.Content != "" {
		note.Content = body.Content
	}
	if err := db.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update note"})
		return
	}
	c.JSON(http.StatusOK, note)
}