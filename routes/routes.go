package routes

import (
	doctor "crud-api/controllers/doctors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartService() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/doctors", doctor.GetDoctors)
		api.GET("/doctors/:id", doctor.GetDoctor)
		api.POST("/doctors", doctor.CreateDoctor)
		api.PUT("/doctors/:id", doctor.UpdateDoctor)
		api.DELETE("/doctors/:id", doctor.DeleteDoctor)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8000")
}
