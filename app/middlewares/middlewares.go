package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userIDSess := session.Get("user_id")
		if userIDSess == nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title":   "Login",
				"message": "",
				"expired": "Yes",
			})
			return
		}
	}
}

func LoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userIDSess := session.Get("user_id")
		if userIDSess != nil {
			c.Redirect(http.StatusFound, "/")
			return
		}
	}
}
