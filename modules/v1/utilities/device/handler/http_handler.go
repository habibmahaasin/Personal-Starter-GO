package handler

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"GuppyTech/pkg/helpers"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (n *deviceHandler) ReceivedData(c *gin.Context) {
	token := "862b34fe2de548cc:cdf66d91b12db8d2"
	getLatestCon, err := n.deviceService.GetLatestCon(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Antares_Device_Id := strings.Replace(getLatestCon.First.Pi, "/antares-cse/cnt-", "", -1)

	c.JSON(http.StatusOK, getLatestCon.First.Con)
}

func (n *deviceHandler) SubscribeWebhook(c *gin.Context) {
	var webhookData models.ObjectAntares1
	if err := c.ShouldBindJSON(&webhookData); err != nil {
		response := helpers.APIRespon("Error, inputan tidak sesuai", 220, "error", nil)
		c.JSON(220, response)
		return
	}
	Antares_Device_Id := strings.Replace(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Pi, "/antares-cse/cnt-", "", -1)
	_, err := n.deviceService.GetDatafromWebhook(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Con, Antares_Device_Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(Antares_Device_Id)
}
