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
	c.SetCookie("Token", token, 3600, "/", "localhost", false, true)

	session.Set("userID", user.User_id)
	session.Set("userName", user.Full_name)
	session.Options(sessions.Options{
		MaxAge: 3600 * 24,
	})
	session.Save()

	c.Redirect(http.StatusFound, "/")
}
