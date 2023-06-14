package view

import (
	"math"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *deviceView) Index(c *gin.Context) {
	var total_devices_off int
	var total_devices_on int
	var average_ph float64
	var counter int
	var average_temperature float64
	var Last_updated_history string
	ListDevice, _ := h.deviceService.GetAllDevices()
	History, _ := h.deviceService.GetDeviceHistory()

	// count total devices
	for i := 0; i < len(ListDevice); i++ {
		if ListDevice[i].Status_name == "Aktif" {
			total_devices_on++
		} else {
			total_devices_off++
		}
	}
	total_devices := len(ListDevice)

	// average ph and temperature
	for i := 0; i < len(History); i++ {
		if History[i].History_ph != 0 && History[i].History_temperature != 0 {
			average_ph += float64(History[i].History_ph)
			average_temperature += float64(History[i].History_temperature)
			counter++
			if counter == 1 {
				Last_updated_history = History[i].History_date_formatter
			}
		}
	}
	average_ph = average_ph / float64(counter)
	average_temperature = average_temperature / float64(counter)

	session := sessions.Default(c)
	user_email := session.Get("email")
	full_name := session.Get("full_name")

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":         "Beranda",
		"email":         user_email,
		"full_name":     full_name,
		"listDevice":    ListDevice,
		"total_devices": total_devices,
		"total_on":      total_devices_on,
		"total_off":     total_devices_off,
		"average_ph":    math.Round(average_ph*100) / 100,
		"average_temp":  math.Round(average_temperature*100) / 100,
		"last_updated":  Last_updated_history,
		"History":       History,
	})
}

func (h *deviceView) ListDevice(c *gin.Context) {
	session := sessions.Default(c)
	user_email := session.Get("email")
	full_name := session.Get("full_name")

	ListDevice, _ := h.deviceService.GetAllDevices()

	c.HTML(http.StatusOK, "list_device.html", gin.H{
		"title":      "Daftar Perangkat",
		"email":      user_email,
		"full_name":  full_name,
		"listDevice": ListDevice,
	})
}

func (h *deviceView) Report(c *gin.Context) {
	session := sessions.Default(c)
	user_email := session.Get("email")
	full_name := session.Get("full_name")
	History, _ := h.deviceService.GetDeviceHistory()

	c.HTML(http.StatusOK, "report.html", gin.H{
		"title":     "Laporan",
		"email":     user_email,
		"full_name": full_name,
		"history":   History,
	})
}
