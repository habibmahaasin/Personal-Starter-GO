package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *deviceView) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Index",
	})
}
