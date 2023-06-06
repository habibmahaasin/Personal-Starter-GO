package view

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *deviceView) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Beranda",
	})
}

func (h *deviceView) ListDevice(c *gin.Context) {
	ListDevice, err := h.deviceService.GetAllDevices()
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "list_device.html", gin.H{
		"title":      "Daftar Perangkat",
		"listDevice": ListDevice,
	})
}

func (h *deviceView) Report(c *gin.Context) {
	History, err := h.deviceService.GetDeviceHistory()
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "report.html", gin.H{
		"title":   "Laporan",
		"history": History,
	})
}
