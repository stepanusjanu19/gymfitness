package model

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Item struct {
	ItemID 		uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"item_id"`
	Name 		string 	   `gorm:"not null;size:255" json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
}

func (item *Item) BeforeCreate(scope *gorm.Scope) error {
	newUUID, err := uuid.NewV4()
	if err!= nil {
        return err
    }

	item.ItemID = newUUID

	return nil
}