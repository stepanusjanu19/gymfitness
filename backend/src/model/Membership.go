package model

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type MembershipType string
type Status string

const (
	Weakly  MembershipType = "weakly"
	Monthly MembershipType = "monthly"
	Yearly  MembershipType = "yearly"

	Active  Status = "active"
	Expired Status = "expired"
)

type Membership struct {
	MembershipID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"membership_id"`
	Category            string         `gorm:"size:255;not null;unique" json:"category"`
	MembershipType      MembershipType `gorm:"type:varchar(10);not null" json:"membership_type"`
	Benefit             []string       `gorm:"type:jsonb" json:"benefit"`                        
	Price               float64        `gorm:"not null" json:"price"`
	Status              Status         `gorm:"type:varchar(10);default:'active';not null" json:"status"`
	IsMembershipTrainer bool           `gorm:"not null;default:false" json:"is_membership_trainer"`
	IsMembershipConsume bool           `gorm:"not null;default:false" json:"is_membership_consume"`
	ItemID              []uuid.UUID    `gorm:"type:uuid[]" json:"item_id"` 
	Description         string         `json:"description"`
	CreatedAt           time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt           *time.Time     `sql:"index" json:"deleted_at"`

	Items []Item `gorm:"many2many:membership_items;" json:"items"`
}

func (membership *Membership) BeforeCreate(scope *gorm.Scope) error {
	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	membership.MembershipID = newUUID

	return nil
}