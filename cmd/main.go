package main

import (
	"mediportal/config"
	"mediportal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"}) // or actual proxy IPs

	config.ConnectDatabase()

	routes.AuthRoutes(r)
	routes.PatientRoutes(r)
	routes.DoctorRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "MediPortal Backend Running!"})
	})

	r.Run()
}
