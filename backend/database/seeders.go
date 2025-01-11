package database

import (
	"fmt"
	"log"
	"backend/src/model"
	"github.com/jinzhu/gorm"
	"backend/lib/helpers"
)

var usersDump = []model.User{
	model.User{
		FirstName: "4DMI1N",
		LastName: "2025",
		Email: "adm1nGym2025@gmail.com",
		Password: "",
		Phone: "123456789",
		Image: "https://cdn-icons-png.flaticon.com/512/6024/6024190.png",
		IsActive: true,
		Role: model.Admin,
	},
	model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@gmail.com",
		Password:  "",
		Phone:     "987654321",
		Image:     "https://cdn-icons-png.flaticon.com/512/6024/6024190.png",
		IsActive:  true,
		Role:      model.Guest,
	},
	model.User{
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@gmail.com",
		Password:  "",
		Phone:     "1122334455",
		Image:     "https://cdn-icons-png.flaticon.com/512/6024/6024190.png",
		IsActive:  true,
		Role:      model.Guest,
	},
}

func LoadSeeders(db *gorm.DB)  {
	if db == nil {
        fmt.Println("Database connection is nil")
        return
    }
	for i, user := range usersDump {
		var password string
		if user.Role == model.Admin {
			password = "4p4j4B01EH"
		} else {
			password = "defaultpassword"
		}
		hashedPassword, err := helpers.HashPassword(password)
		if err != nil {
			log.Fatalf("Error hashing password: %v", err)
		}
		usersDump[i].Password = hashedPassword
		err = db.Debug().Create(&user).Error
		if err != nil {
			log.Fatalf("cannot seed data: %v", err)
		}
	}
	log.Println("Seeding completed!")
}