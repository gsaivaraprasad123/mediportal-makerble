package routes

import (
	"mediportal/controllers"
	"mediportal/middleware"

	"github.com/gin-gonic/gin"
)

func DoctorRoutes(r *gin.Engine) {
	d := r.Group("/api/doctor/patients")
	d.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("doctor"))

	d.GET("/", controllers.DoctorGetAllPatients)
	d.PUT("/:id", controllers.DoctorUpdatePatients)
}
