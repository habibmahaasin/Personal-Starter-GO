package database

import (
	"Batumbuah/app/config"
	"Batumbuah/modules/v1/utilities/user/models"
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
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.CheckInLog{})
	Db.AutoMigrate(&models.UserStats{})
	Db.AutoMigrate(&models.TestInformation{})
	return Db
}
