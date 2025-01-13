package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB  {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", 
		viper.GetString("DATABASE_HOST"),
		viper.GetString("DATABASE_PORT"),
		viper.GetString("DATABASE_USERNAME"),
		viper.GetString("DATABASE_NAME"),
		viper.GetString("DATABASE_PASSWORD"),
	)

	var err error

	db, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Could not connect to database :%v", err)
	}

	DB = db

	log.Println("Database Connection established succcessfully")
	return db
}

