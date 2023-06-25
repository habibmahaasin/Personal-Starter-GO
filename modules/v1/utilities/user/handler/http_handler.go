package handler

import (
	"GuppyTech/modules/v1/utilities/user/models"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (n *userHandler) Login(c *gin.Context) {
	session := sessions.Default(c)
	var input models.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		return
	}

	user, err := n.userService.Login(input)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":   "Login",
			"message": "Username/ Password yang anda masukkan Salah!",
		})
		return
	}

	token, _ := n.jwtoken.GenerateToken(user.User_id, user.Full_name, user.Role_id)
	c.SetCookie("Token", token, 21600, "/", "guppy.tech", false, true)

	session.Set("email", user.Email)
	session.Set("full_name", user.Full_name)
	session.Set("user_id", user.User_id)
	session.Options(sessions.Options{
		MaxAge: 3600 * 6,
	})
	session.Save()

	c.Redirect(http.StatusFound, "/")
}

func (n *userHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "GuppyTech",
		MaxAge: -1,
	})

	c.Redirect(http.StatusFound, "/login")
}
