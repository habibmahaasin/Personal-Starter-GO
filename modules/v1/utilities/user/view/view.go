package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *userView) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title":   "Login",
		"message": "",
	})
}
