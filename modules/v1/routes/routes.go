package routes

import (
	"GuppyTech/app/config"
	deviceHandlerV1 "GuppyTech/modules/v1/utilities/device/handler"
	deviceviewV1 "GuppyTech/modules/v1/utilities/device/view"
	basic "GuppyTech/pkg/basic_auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParseTmpl(router *gin.Engine) *gin.Engine { //Load HTML Template
	router.Static("/assets", "./public/assets")
	router.Static("/img", "./public/assets/img")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/vendor", "./public/assets/vendor")
	return router
}

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	deviceHandlerV1 := deviceHandlerV1.Handler(db)
	deviceViewV1 := deviceviewV1.View(db)

	// Routing Website Service
	device := router.Group("/", basic.Auth(conf))
	device.GET("/", deviceViewV1.Index)
	device.GET("/login", deviceViewV1.Login)
	device.GET("/daftar-perangkat", deviceViewV1.ListDevice)

	//Routing API Service
	api := router.Group("/api/v1")
	api.GET("/antares-data", deviceHandlerV1.ReceivedData)
	api.POST("/webhook", deviceHandlerV1.SubscribeWebhook)
	router = ParseTmpl(router)
	return router
}
