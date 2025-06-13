package routes

import (
	"mediportal/controllers"
	"mediportal/middleware"

	"github.com/gin-gonic/gin"
)

func PatientRoutes(r *gin.Engine) {
	p := r.Group("/api/patients")
	p.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("receptionist"))

	p.POST("/", controllers.CreatePatient)
	p.GET("/", controllers.GetAllPatients)
	p.GET("/:id", controllers.GetPatientByID)
	p.PUT("/:id", controllers.UpdatePatient)
	p.DELETE("/:id", controllers.DeletePatient)
}
