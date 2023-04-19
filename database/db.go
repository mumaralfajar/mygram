package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"mygram/models"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	var (
		host     = viper.Get("PGHOST")
		port     = viper.Get("PGPORT")
		user     = viper.Get("PGUSER")
		dbname   = viper.Get("PGDATABASE")
		password = viper.Get("PGPASSWORD")
	)

	config := fmt.Sprintf("host=%v user=%v password=%v port=%v dbname=%v sslmode=disable", host, user, password, port, dbname)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	log.Println("connected to database successfully")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
