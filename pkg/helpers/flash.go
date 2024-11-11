package helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetFlashMessage(c *gin.Context, status string, message string) {
    session := sessions.Default(c)
    session.Set("status", status)
    session.Set("message", message)
    session.Save()
}

func GetAndClearFlashMessage(c *gin.Context) (status, message string) {
    session := sessions.Default(c)
    status, _ = session.Get("status").(string)
    message, _ = session.Get("message").(string)
    session.Delete("status")
    session.Delete("message")
    session.Save()
    return status, message
}