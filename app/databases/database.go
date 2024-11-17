package database

import (
	"Go_Starter/app/config"
	"Go_Starter/modules/v1/utilities/user/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(conf config.Conf) *gorm.DB {
	var dsn string
	var Db *gorm.DB
	var err error

	switch conf.Db.Type {
	case "postgres":
		dsn = "host=" + conf.Db.Host + " user=" + conf.Db.User + " password=" + conf.Db.Pass + " dbname=" + conf.Db.Name + " port=" + conf.Db.Port + " TimeZone=Asia/Jakarta"
		Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		// DSN format for MySQL: "user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		dsn = conf.Db.User + ":" + conf.Db.Pass + "@tcp(" + conf.Db.Host + ":" + conf.Db.Port + ")/" + conf.Db.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		log.Fatalln("Unsupported database type. Please use 'postgres' or 'mysql'.")
	}

	if err != nil {
		log.Fatalln("Failed to connect to database:", err.Error())
	}

	Db.AutoMigrate(&models.User{})
	return Db
}
