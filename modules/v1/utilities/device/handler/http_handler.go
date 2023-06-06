package handler

import (
	"GuppyTech/modules/v1/utilities/device/models"
	api "GuppyTech/pkg/api_response"
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
		response := helpers.APIRespon("Error, Format Input Tidak Sesuai", 220, "error", nil)
		c.JSON(220, response)
		return
	}
	Antares_Device_Id := strings.Replace(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Pi, "/antares-cse/cnt-", "", -1)
	Input, err := n.deviceService.GetDatafromWebhook(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Con, Antares_Device_Id)
	if err != nil {
		response := api.APIRespon("Error, Please Check "+err.Error(), 500, "error", nil)
		c.JSON(500, response)
		return
	} else {
		response := api.APIRespon("Success", 200, "success", Input)
		c.JSON(200, response)
		return
	}
}
