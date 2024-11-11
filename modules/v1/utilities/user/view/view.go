package view

import (
	"Batumbuah/pkg/helpers"
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
    userID := session.Get("user_id")
    checkInLogs, _ := h.userService.GetCheckInLogs(userID.(string))
    userStats, _ := h.userService.GetUserStats(userID.(string))
    status, message := helpers.GetAndClearFlashMessage(c)

    c.HTML(http.StatusOK, "index.html", gin.H{
        "title":       "Index",
        "status":      status,
        "message":     message,
        "email":       email,
        "name":        name,
        "checkInLogs": checkInLogs,
        "userStats":   userStats,
    })
}

func (h *userView) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title":   "Register",
	})
}