package model

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Role string

const (
	Admin   Role = "admin"
	Trainer Role = "trainer"
	Member  Role = "member"
	Guest   Role = "Guest"
)

var ValidRoles = []string{string(Admin), string(Trainer), string(Member), string(Guest)}

type User struct {
	UserId           uuid.UUID          `gorm:"type:uuid; primary_key;default:uuid_generate_v4()" json:"user_id"`
	FirstName        string             `gorm:"size:255;not null" json:"firstname"`
	LastName         string             `gorm:"size:255" json:"lastname"`
	FullName         string             `gorm:"-" json:"fullname"`
	Email            string             `gorm:"unique; not null" json:"email"`
	Password         string             `gorm:"not null" json:"password"`
	Phone            string             `json:"phone"`
	Image            string             `json:"image"`
	IsActive         bool               `gorm:"default:false" json:"is_active"`
	Role             Role               `gorm:"type:role_user; default: 'guest'" json:"role"`
	CreatedAt        time.Time          `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt         time.Time          `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	VerificationUser []VerificationUser `gorm:"foreignkey:UserId;constraint:onDelete:CASCADE;" json:"verification_users"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	user.UserId = newUUID
	return nil
}

func IsValidRoles(role string) bool {
	for _, validRole := range ValidRoles {
		if role == validRole {
			return true
		}
	}
	return false
}
