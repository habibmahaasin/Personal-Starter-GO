package handler

import (
	"GuppyTech/modules/v1/utilities/device/models"
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

	fmt.Println(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Con)
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
	antares_id := c.Param("antares")
	power := c.Param("power")
	token := "f784524323f73064:4c0b580400028426"

	err := h.deviceService.Control(id, power, mode)
	if err != nil {
		fmt.Println(err)
		return
	}

	session := sessions.Default(c)
	user_id := session.Get("user_id").(string)
	getDeviceById, _ := h.deviceService.GetDeviceById(user_id, id)

	for i := 0; i < 2; i++ {
		err = h.deviceService.PostControlAntares(antares_id, token, power, mode, getDeviceById.Ph_calibration_firstval, getDeviceById.Ph_calibration_secval)
		time.Sleep(2 * time.Second)
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
	session := sessions.Default(c)
	user_id := session.Get("user_id").(string)

	err := c.ShouldBind(&input)
	device, _ := h.deviceService.GetDeviceById(user_id, input.Device_id)

	for i := 0; i < 2; i++ {
		if device.Status_id == "10" {
			device.Status_id = "0"
		} else if device.Status_id == "11" {
			device.Status_id = "1"
		}
		err = h.deviceService.PostControlAntares(device.Antares_id, token, device.Status_id, device.Mode_id, input.Ph_calibration_firstval, input.Ph_calibration_secval)
		time.Sleep(2 * time.Second)
	}

	err = h.deviceService.Calibration(input)
	if err != nil {
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound, "/kalibrasi-sensor")
}
