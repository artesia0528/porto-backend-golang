package handlers

import (
	"net/http"
	"portfolio-backend/internal/database"
	"portfolio-backend/internal/models"
	"github.com/gin-gonic/gin"
)

// Publik: siapa saja bisa lihat
func GetProjects(c *gin.Context) {
	var projects []models.Project
	database.DB.Find(&projects)
	c.JSON(http.StatusOK, projects)
}

// Butuh login: hanya admin bisa tambah project
func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	database.DB.Create(&project)
	c.JSON(http.StatusCreated, project)
}

func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Project{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Project dihapus"})
}