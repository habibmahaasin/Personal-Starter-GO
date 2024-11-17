package handler

import (
	"Go_Starter/modules/v1/utilities/user/models"
	api "Go_Starter/pkg/api_response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) Register(c *gin.Context) {
	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, api.APIRespon("Invalid input", http.StatusBadRequest, "error", gin.H{
			"error": err.Error(),
		}))
		return
	}

	err := h.userService.Register(input.FullName, input.Email, input.Password, input.Address, 2)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIRespon("Registration failed", http.StatusBadRequest, "error", gin.H{
			"error": err.Error(),
		}))
		return
	}

	c.JSON(http.StatusOK, api.APIRespon("Registration successful, please log in", http.StatusOK, "success", nil))
}

func (h *userHandler) Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, api.APIRespon("Invalid input", http.StatusBadRequest, "error", gin.H{
			"error": err.Error(),
		}))
		return
	}

	user, err := h.userService.Login(input)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, api.APIRespon("Invalid username or password", http.StatusUnauthorized, "error", nil))
		return
	}

	token, err := h.jwtoken.GenerateToken(user.UserID, user.FullName, int(user.RoleID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, api.APIRespon("Failed to generate token", http.StatusInternalServerError, "error", nil))
		return
	}

	// Set the token in a cookie for monolithic apps
	// session := sessions.Default(c)
	// c.SetCookie("Token", token, 21600, "/", "localhost", false, true)
	// session.Set("email", user.Email)
	// session.Set("full_name", user.FullName)
	// session.Set("user_id", user.UserID)
	// session.Options(sessions.Options{
	// 	MaxAge: 3600 * 6, // 6 hours
	// })
	// session.Save()

	c.JSON(http.StatusOK, api.APIRespon("Login successful", http.StatusOK, "success", gin.H{
		"user": gin.H{
			"user_id":   user.UserID,
			"full_name": user.FullName,
			"email":     user.Email,
		},
		"token": token,
	}))
}