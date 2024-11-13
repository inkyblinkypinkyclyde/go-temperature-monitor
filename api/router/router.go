package router

import (
	"net/http"
	"time"

	"main/models"

	"github.com/gin-gonic/gin"
)

type DefaultReceiver struct {
}

func (r *DefaultReceiver) LogTemperature(c *gin.Context) {
	var temperatureReport models.TemperatureReport
	if err := c.ShouldBindJSON(&temperatureReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad iot default request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "recieved iot request",
	})
	temperatureReport.RcvdTime = time.Now()
	// log temperature request here
}
