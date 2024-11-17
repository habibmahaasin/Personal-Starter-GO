package routes

import (
	"Go_Starter/app/config"
	userHandlerV1 "Go_Starter/modules/v1/utilities/user/handler"
	userViewV1 "Go_Starter/modules/v1/utilities/user/view"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParseTmpl(router *gin.Engine) *gin.Engine {
	router.Static("/assets", "./public/assets")
	router.Static("/img", "./public/assets/img")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/vendor", "./public/assets/vendor")
	return router
}

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	userHandlerV1 := userHandlerV1.Handler(db)
	userViewV1 := userViewV1.View(db)

	// Routing API Service
	api := router.Group("/api/v1")
	api.POST("/register", userHandlerV1.Register)
	api.POST("/login", userHandlerV1.Login)
	api.GET("/logout", userHandlerV1.Logout)

	// Routing View
	view := router.Group("/")
	view.GET("/", userViewV1.Index)

	router = ParseTmpl(router)
	return router
}
