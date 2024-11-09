package view

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *userView) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title":   "Login",
		"message": "",
	})
}

func (h *userView) Index(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get("email")
	name := session.Get("full_name")

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Index",
		// "message": user,
		"email": email,
		"name":  name,
	})
}