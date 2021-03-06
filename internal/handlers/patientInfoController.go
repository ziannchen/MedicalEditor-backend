package handlers

import (
	"editor-backend/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPatientInfoList(c *gin.Context) {
	patientInfos, err := services.GetPatientInfos()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"patientInfos": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"patientInfos": patientInfos,
		})
	}
}

func GetPatientInfo(c *gin.Context) {
	patientId := c.Query("patientId")
	patientInfo, err := services.GetPatientInfoByPatientId(patientId)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"patientInfo": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"patientInfo": patientInfo,
		})
	}
}
