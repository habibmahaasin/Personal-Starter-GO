package handler

import (
	"GuppyTech/modules/v1/utilities/device/models"
	api "GuppyTech/pkg/api_response"
	"GuppyTech/pkg/helpers"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	fmt.Println("data : ", webhookData.First.M2m_nev.M2m_rep.M2m_cin.Con)
	Antares_Device_Id := strings.Replace(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Pi, "/antares-cse/cnt-", "", -1)
	_, err := h.deviceService.GetDatafromWebhook(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Con, Antares_Device_Id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (h *deviceHandler) Control(c *gin.Context) {
	page := c.Param("page")
	id := c.Param("id")
	mode := c.Param("mode")
	power := c.Param("power")
	token := "f784524323f73064:4c0b580400028426"
	session := sessions.Default(c)
	user_id := session.Get("user_id").(string)

	getData, _ := h.deviceService.GetDeviceById(user_id, id)
	err := h.deviceService.PostControlAntares(getData.Antares_id, token, power, mode)
	err = h.deviceService.Control(id, power, mode)
	if err != nil {
		fmt.Println(err)
		return
	}

	if page == "detail_perangkat" {
		c.Redirect(http.StatusFound, "/detail-perangkat/"+id)
	} else if page == "daftar_perangkat" {
		c.Redirect(http.StatusFound, "/daftar-perangkat")
	}
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

func (h *deviceHandler) DeleteDevice(c *gin.Context) {
	device_id := c.Param("id")
	err := h.deviceService.DeleteDevice(device_id)
	if err != nil {
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound, "/daftar-perangkat")
}

func (h *deviceHandler) EditDevice(c *gin.Context) {
	device_id := c.Param("id")
	var input models.DeviceInput

	err := c.ShouldBind(&input)

	err = h.deviceService.UpdateDeviceById(input, device_id)
	if err != nil {
		log.Println(err)
		return
	}

	c.Redirect(http.StatusFound, "/daftar-perangkat")
}

func (h *deviceHandler) Calibration(c *gin.Context) {
	var input models.PhCalibration
	token := "f784524323f73064:4c0b580400028426"

	err := c.ShouldBind(&input)
	err = h.deviceService.PostCalibrationAntares(token, input)

	err = h.deviceService.Calibration(input)
	if err != nil {
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound, "/kalibrasi-sensor")
}

func (h *deviceHandler) APIControlling(c *gin.Context) {
	var input models.ControllingAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.APIRespon("Error, Format Input Tidak Sesuai", 220, "error", nil)
		c.JSON(220, response)
		return
	}

	token := "f784524323f73064:4c0b580400028426"

	getData, _ := h.deviceService.GetDeviceById(input.User_id, input.Device_id)
	err = h.deviceService.PostControlAntares(getData.Antares_id, token, input.Power, input.Mode)
	err = h.deviceService.Control(input.Device_id, input.Power, input.Mode)

	if err == nil {
		response := api.APIRespon("Success", 200, "success", input)
		c.JSON(200, response)
	}
}
