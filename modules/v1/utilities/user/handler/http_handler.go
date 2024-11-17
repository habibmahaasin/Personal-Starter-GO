package handler

import (
	"Go_Starter/modules/v1/utilities/user/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *userHandler) Register(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// Register the user
	err := h.userService.Register(input.FullName, input.Email, input.Password, input.Address, 2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Registration successful, please log in",
	})
}

func (h *userHandler) Login(c *gin.Context) {
	session := sessions.Default(c)
	var input models.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		return
	}

	user, err := h.userService.Login(input)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":   "Login",
			"message": "Username/ Password yang anda masukkan Salah!",
		})
		return
	}

	token, _ := h.jwtoken.GenerateToken(user.UserID, user.FullName, int(user.RoleID))
	fmt.Println(token)
	c.SetCookie("Token", token, 21600, "/", "localhost", false, true)

	session.Set("email", user.Email)
	session.Set("full_name", user.FullName)
	session.Set("user_id", user.UserID)
	session.Options(sessions.Options{
		MaxAge: 3600 * 6,
	})
	session.Save()

	c.Redirect(http.StatusFound, "/")
}

func (h *userHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "Go_Starter",
		MaxAge: -1,
	})

	c.Redirect(http.StatusFound, "/login")
}