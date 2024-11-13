package router

import (
	"fmt"
	"net/http"
	"time"

	models "github.com/inkyblinkypinkyclyde/go-home/api/models"

	"github.com/gin-gonic/gin"
)

type DefaultReceiver struct {
}

func (r *DefaultReceiver) HealthCheck(c *gin.Context) {
	var areYouThereReq models.AreYouThere
	if err := c.ShouldBindJSON(&areYouThereReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad are you there request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("I am here, just recieved '%s'", areYouThereReq.Phrase),
	})
}

func (r *DefaultReceiver) LogTemperature(c *gin.Context) {
	var temperatureReport models.NodeReport
	if err := c.ShouldBindJSON(&temperatureReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad iot default request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "recieved iot request",
	})
	temperatureReport.RcvdTime.Time = time.Now()
	temperatureReport.RcvdTime.Valid = true
	// log temperature request here
}
