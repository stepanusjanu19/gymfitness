package database

import (
	"backend/src/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func LoadMigrate(db *gorm.DB)  {
	if db == nil {
        fmt.Println("Database connection is nil")
        return
    }
	err := db.Debug().DropTableIfExists(&model.User{}, &model.VerificationUser{}, &model.Membership{}, &model.Item{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&model.User{}, &model.VerificationUser{}, &model.Membership{}, &model.Item{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	log.Println("Migration completed!")
}