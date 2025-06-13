package controllers

import (
	"mediportal/config"
	"mediportal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	patient.ID = uuid.New()

	if err := config.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Patient created successfully", "patient": patient})
}

func GetAllPatients(c *gin.Context) {
	var patients []models.Patient
	if err := config.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patients"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"patients": patients})
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := config.DB.First(&patient, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var input models.Patient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	patient.Name = input.Name
	patient.Age = input.Age
	patient.Gender = input.Gender
	patient.Condition = input.Condition

	if err := config.DB.Save(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully", "patient": patient})
}

func DeletePatient(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Patient{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}
