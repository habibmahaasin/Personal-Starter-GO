package handler

import (
	"GuppyTech/modules/v1/utilities/device/models"
	api "GuppyTech/pkg/api_response"
	"GuppyTech/pkg/helpers"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *deviceHandler) SubscribeWebhook(c *gin.Context) {
	var webhookData models.ObjectAntares1
	if err := c.ShouldBindJSON(&webhookData); err != nil {
		response := helpers.APIRespon("Error, Format Input Tidak Sesuai", 220, "error", nil)
		c.JSON(220, response)
		return
	}
	Antares_Device_Id := strings.Replace(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Pi, "/antares-cse/cnt-", "", -1)
	Input, err := h.deviceService.GetDatafromWebhook(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Con, Antares_Device_Id)
	if err == nil {
		response := api.APIRespon("Success", 200, "success", Input)
		c.JSON(200, response)
		return
	}
}

func (h *deviceHandler) Control(c *gin.Context) {
	id := c.Param("id")
	mode := c.Param("mode")
	antares_id := c.Param("antares")
	power := c.Param("power")
	token := "862b34fe2de548cc:cdf66d91b12db8d2"

	err := h.deviceService.Control(id, power, mode)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 2; i++ {
		h.deviceService.PostControlAntares(antares_id, token, power, mode)
		time.Sleep(2 * time.Second)
	}

	c.Redirect(http.StatusFound, "/daftar-perangkat")
	return
}

func (h *deviceHandler) AddDevice(c *gin.Context) {
	session := sessions.Default(c)
	var input models.DeviceInput

	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		return
	}
	user_id := session.Get("user_id").(string)
	err = h.deviceService.AddDevice(input, user_id)

	c.Redirect(http.StatusFound, "/daftar-perangkat")
}
