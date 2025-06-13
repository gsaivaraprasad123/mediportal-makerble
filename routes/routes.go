package routes

import (
	"mediportal/controllers"
	"mediportal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	public := r.Group("/api")
	{
		public.POST("/signup", controllers.Signup)
		public.POST("/login", controllers.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/patients", controllers.CreatePatient)
		protected.GET("/patients", controllers.GetAllPatients)
		protected.PUT("/patients/:id", controllers.UpdatePatient)
		protected.DELETE("/patients/:id", controllers.DeletePatient)

		protected.GET("/doctor/patients", controllers.DoctorGetAllPatients)
		protected.PUT("/doctor/patients/:id", controllers.DoctorUpdatePatients)
	}
}
