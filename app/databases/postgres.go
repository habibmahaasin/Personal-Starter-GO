package database

import (
	"GuppyTech/app/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(conf config.Conf) *gorm.DB {

	dsn := "host=" + conf.Db.Host + " user=" + conf.Db.User + " password=" + conf.Db.Pass + " dbname=" + conf.Db.Name + " port=" + conf.Db.Port + " TimeZone=Asia/Jakarta"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	// digunakan untuk auto migrate dari entity ke database
	// db.AutoMigrate(&models.User{})

	return Db
}
