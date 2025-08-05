package database

import (
	"api/src/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func LoadMigrate(db *gorm.DB) {
	if db == nil {
		fmt.Println("Database connection is nil")
		return
	}

	err := db.Exec(`
		DO $$ BEGIN
			CREATE TYPE role_user AS ENUM ('guest', 'admin', 'trainer');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;
	`).Error
	if err != nil {
		log.Fatalf("Failed to create enum type: %v", err)
	}

	err = db.Debug().DropTableIfExists(&model.User{}, &model.VerificationUser{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&model.User{}, &model.VerificationUser{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	log.Println("Migration completed!")
}
