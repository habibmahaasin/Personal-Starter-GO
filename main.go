package main

import (
	"Go_Starter/app/config"
	database "Go_Starter/app/databases"
	routesV1 "Go_Starter/modules/v1/routes"
	"Go_Starter/pkg/html"
	error "Go_Starter/pkg/http-error"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setup() (*gorm.DB, config.Conf, *gin.Engine) {
	conf, err := config.Init()
	gin.SetMode(conf.App.Mode)
	if err != nil {
		log.Fatal(err)
	}
	db := database.Init(conf)

	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(conf.App.Secret_key))
	router.Use(sessions.Sessions("Go_Starter", cookieStore))
	router.HTMLRender = html.Render("./public/templates")

	//Error Handling for 404 Not Found Page and Method Not Allowed
	router.NoRoute(error.PageNotFound())
	router.NoMethod(error.NoMethod())
	return db, conf, router
}

func main() {
	_, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	router := routesV1.Init(setup()) //Version 1
	router.Run()
	// router.Run(":" + conf.App.Port)
}
