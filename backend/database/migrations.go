package database

import (
	"log"
	"backend/src/model"
	"github.com/jinzhu/gorm"
)

func LoadMigrate(db *gorm.DB)  {
	err := db.Debug().DropTableIfExists(&model.User{}, &model.VerificationUser{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&model.User{}, &model.VerificationUser{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	log.Println("Migration completed!")
}