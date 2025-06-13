package controllers

import (
	"mediportal/config"
	"mediportal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoctorGetAllPatients(c *gin.Context) {
	var patients []models.Patient
	if err := config.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func DoctorUpdatePatients(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := config.DB.First(&patient, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var input struct {
		Condition string `json:"condition"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update only condition
	patient.Condition = input.Condition
	patient.UpdatedBy = c.GetString("email")

	if err := config.DB.Save(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, patient)
}
