package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func LoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func CreatePatientHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Patient created successfully",
	})
}
