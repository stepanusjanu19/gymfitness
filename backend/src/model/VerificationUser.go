package model

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type VerificationUser struct {
	VerificationId uuid.UUID `gorm:"type:uuid; primary_key;default:uuid_generate_v4()" json:"verification_id"`
	UserId         uuid.UUID `gorm:"type:uuid; not null" json:"user_id"`
	Otp            int       `json:"otp"`
	IsUsed         bool      `gorm:"default:false" json:"is_used"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	User           User      `gorm:"foreignkey:UserId;constraint:onDelete:CASCADE;" json:"user"`
}

func (verificationUser *VerificationUser) BeforeCreate(scope *gorm.Scope) error {
	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	verificationUser.VerificationId = newUUID
	return nil
}
