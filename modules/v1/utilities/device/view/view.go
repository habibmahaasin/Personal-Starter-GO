package view

import (
	"encoding/json"
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

	session := sessions.Default(c)
	user_email := session.Get("email")
	full_name := session.Get("full_name")
	user_id := session.Get("user_id")

	ListDevice, _ := h.deviceService.GetAllDevices(user_id.(string))
	History, GraphHistory, _ := h.deviceService.GetDeviceHistory(user_id.(string))

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
	if math.IsNaN(average_ph) {
		average_ph = 0
	}
	average_temperature = average_temperature / float64(counter)
	if math.IsNaN(average_temperature) {
		average_temperature = 0
	}

	convJsonHistory, _ := json.Marshal(GraphHistory)
	var JSONHistory interface{}
	json.Unmarshal(convJsonHistory, &JSONHistory)

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
		"History":       JSONHistory,
	})
}

func (h *deviceView) ListDevice(c *gin.Context) {
	session := sessions.Default(c)
	user_email := session.Get("email")
	full_name := session.Get("full_name")
	user_id := session.Get("user_id")

	ListDevice, _ := h.deviceService.GetAllDevices(user_id.(string))

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
	user_id := session.Get("user_id")

	History, _, _ := h.deviceService.GetDeviceHistory(user_id.(string))

	c.HTML(http.StatusOK, "report.html", gin.H{
		"title":     "Laporan",
		"email":     user_email,
		"full_name": full_name,
		"history":   History,
	})
}

func (h *deviceView) AddDevice(c *gin.Context) {
	session := sessions.Default(c)
	user_email := session.Get("email")
	full_name := session.Get("full_name")
	brandList, _ := h.deviceService.GetDeviceBrands()

	c.HTML(http.StatusOK, "add_device.html", gin.H{
		"title":     "Tambah Perangkat",
		"email":     user_email,
		"full_name": full_name,
		"brands":    brandList,
	})
}

func (h *deviceView) DetailDevice(c *gin.Context) {
	d_id := c.Param("id")
	session := sessions.Default(c)
	user_email := session.Get("email")
	full_name := session.Get("full_name")
	u_id := session.Get("user_id")

	DetailDevice, err := h.deviceService.GetDeviceById(u_id.(string), d_id)
	History, GraphHistory, _ := h.deviceService.GetDeviceHistoryById(d_id, u_id.(string))

	// checking access
	if err != nil || DetailDevice.Device_id == "" {
		c.HTML(404, "error_404.html", gin.H{
			"title": "Page Not Found",
		})
		return
	}

	var latest_ph_value float32
	var latest_temperature_value float32
	var latest_dissolved_oxygen_value float32

	for i := 0; i < len(History); i++ {
		if History[i].History_ph != 0 && History[i].History_temperature != 0 && History[i].History_dissolved_oxygen != 0 {
			latest_ph_value = History[i].History_ph
			latest_temperature_value = History[i].History_temperature
			latest_dissolved_oxygen_value = History[i].History_dissolved_oxygen
			break
		}
	}

	// convert graph history to json
	convJsonHistory, _ := json.Marshal(GraphHistory)
	var JSONHistory interface{}
	json.Unmarshal(convJsonHistory, &JSONHistory)

	c.HTML(http.StatusOK, "detail_device.html", gin.H{
		"title":        "Detail Perangkat",
		"email":        user_email,
		"full_name":    full_name,
		"latitude":     DetailDevice.Latitude,
		"longitude":    DetailDevice.Longitude,
		"deviceInfo":   DetailDevice,
		"graphHistory": JSONHistory,
		"latest_ph":    latest_ph_value,
		"latest_temp":  latest_temperature_value,
		"latest_do":    latest_dissolved_oxygen_value,
		"history":      History,
	})
}

func (h *deviceView) EditDevice(c *gin.Context) {
	session := sessions.Default(c)
	user_email := session.Get("email")
	full_name := session.Get("full_name")
	u_id := session.Get("user_id")
	d_id := c.Param("id")

	brandList, _ := h.deviceService.GetDeviceBrands()
	DetailDevice, _ := h.deviceService.GetDeviceById(u_id.(string), d_id)

	c.HTML(http.StatusOK, "edit_device.html", gin.H{
		"title":          "Ubah Perangkat",
		"email":          user_email,
		"full_name":      full_name,
		"selectedDevice": DetailDevice,
		"brands":         brandList,
	})
}

func (h *deviceView) Calibration(c *gin.Context) {
	session := sessions.Default(c)
	user_email := session.Get("email")
	full_name := session.Get("full_name")
	user_id := session.Get("user_id")
	ListDevice, _ := h.deviceService.GetAllDevices(user_id.(string))

	c.HTML(http.StatusOK, "calibration.html", gin.H{
		"title":      "Kalibrasi Sensor",
		"email":      user_email,
		"full_name":  full_name,
		"listDevice": ListDevice,
	})
}
