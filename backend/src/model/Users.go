package model

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Role string 

const (
	Admin Role = "admin"
	Trainer Role = "trainer"
	Member Role = "member"
	Guest Role = "Guest"
)

var ValidRoles = []string{string(Admin), string(Trainer), string(Member), string(Guest)}

type User struct {
	gorm.Model

	UserId uuid.UUID `gorm:"type:uuid; primary_key;default:uuid_generate_v4()" json:"user_id"`
	FullName string `gorm:"not null" json:"fullname"`
	Email string  `gorm:"unique; not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Phone string `json:"phone"`
	Image string `json:"image"`
	IsActive bool `gorm:"default:false" json:"is_active"`
	Role Role `gorm:"type:role_user; default: 'guest'" json:"role"`		

}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	newUUID, err := uuid.NewV4()

	if err != nil {
		return err
	}

	user.UserId = newUUID
	return nil
}

func IsValidRoles(role string) bool  {
	for _, validRole := range ValidRoles {
		if role == validRole {
			return true
		}
	}

	return false
}
