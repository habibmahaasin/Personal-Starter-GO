package routes

import (
	"GuppyTech/app/config"
	"GuppyTech/app/middlewares"
	deviceHandlerV1 "GuppyTech/modules/v1/utilities/device/handler"
	deviceviewV1 "GuppyTech/modules/v1/utilities/device/view"
	userHandlerV1 "GuppyTech/modules/v1/utilities/user/handler"
	userViewV1 "GuppyTech/modules/v1/utilities/user/view"

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
	deviceHandlerV1 := deviceHandlerV1.Handler(db, conf)
	deviceViewV1 := deviceviewV1.View(db, conf)
	userHandlerV1 := userHandlerV1.Handler(db)
	userViewV1 := userViewV1.View(db)

	//Routing API Service
	api := router.Group("/api/v1")
	api.POST("/webhook", deviceHandlerV1.SubscribeWebhook)
	api.POST("/device-controll", deviceHandlerV1.APIControlling)

	// Routing Website Service
	device := router.Group("/")
	device.GET("/login", middlewares.LoggedIn(), userViewV1.Login)
	device.POST("/login", userHandlerV1.Login)
	device.GET("/logout", userHandlerV1.Logout)

	device.GET("/", middlewares.IsLogin(), deviceViewV1.Index)
	device.GET("/daftar-perangkat", middlewares.IsLogin(), deviceViewV1.ListDevice)
	device.GET("/laporan", middlewares.IsLogin(), deviceViewV1.Report)
	device.GET("/tambah-perangkat", middlewares.IsLogin(), deviceViewV1.AddDevice)
	device.POST("/tambah-perangkat", middlewares.IsLogin(), deviceHandlerV1.AddDevice)
	device.GET("/detail-perangkat/:id", middlewares.IsLogin(), deviceViewV1.DetailDevice)
	device.GET("/control/:page/:id/:mode/:power", middlewares.IsLogin(), deviceHandlerV1.Control)
	device.GET("/hapus-perangkat/:id", middlewares.IsLogin(), deviceHandlerV1.DeleteDevice)
	device.GET("/ubah-perangkat/:id", middlewares.IsLogin(), deviceViewV1.EditDevice)
	device.POST("/ubah-perangkat/:id", middlewares.IsLogin(), deviceHandlerV1.EditDevice)
	device.GET("/kalibrasi-sensor", middlewares.IsLogin(), deviceViewV1.Calibration)
	device.POST("/kalibrasi-sensor", middlewares.IsLogin(), deviceHandlerV1.Calibration)
	router = ParseTmpl(router)
	return router
}
