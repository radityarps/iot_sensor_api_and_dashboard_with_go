package routes

import (
	"iot_api_with_go/controllers"

	"github.com/gin-gonic/gin"
)

func SensorStatusRoutes(router *gin.Engine) {
	router.GET("/", controllers.GetStatusSensor)
	router.PUT("/sensor/:id", controllers.UpdateStatusSensor)
	router.POST("/sensor", controllers.AddDataSensor)
	router.PUT("/sensor/update-all", controllers.UpdateAllStatusSensor)
	router.GET("/value-sensor", controllers.GetOnlyValueSensor)
}
