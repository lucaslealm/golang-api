package routes

import (
	doctor "crud-api/controllers/doctors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	version := router.Group("/v1")
	{
		version.GET("/doctors", doctor.GetDoctors)
		version.GET("/doctors/:id", doctor.GetDoctor)
		version.POST("/doctors", doctor.CreateDoctor)
		version.PUT("/doctors/:id", doctor.UpdateDoctor)
		version.DELETE("/doctors/:id", doctor.DeleteDoctor)
	}
	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})
	return router
}
