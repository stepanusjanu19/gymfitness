package services

import (
	"backend/domain/requests"
	"backend/lib/config"
	"backend/src/model"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)



func FindItemAll() ([]model.Item, error) {
	DB := config.DB
	
	var items []model.Item
    err := DB.Find(&items).Error
    if err != nil {
        return nil, err
    }

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []model.Item{}, nil
	}

    return items, nil
}

func FindByItemID(itemID string) (model.Item, error) {
	var item model.Item
	DB := config.DB
	
	ItemUUID, err := uuid.FromString(itemID)
	if err!= nil {
		fmt.Println(err.Error())
        return item, fmt.Errorf("Error parsing item ID")
    }
	if err := DB.Where("item_id = ?", ItemUUID).First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return item, fmt.Errorf("Item not found")
		}

		return item, fmt.Errorf("Error find item")
	}

	return item, nil
}

func CreateItem(request requests.SaveItem) (model.Item, error) {
	var item model.Item
	DB := config.DB
	
	if err := DB.Where("name = ?", request.Name).First(&item).Error; err == nil {
		return item, fmt.Errorf("Name item already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return item, fmt.Errorf("Error finding item by name")
	}

	newItem := model.Item{
		Name: request.Name,
		Description: request.Description,
	}

    if err := DB.Create(&newItem).Error; err!= nil {
		fmt.Println(err.Error())
        return newItem, fmt.Errorf("Error save item")
    }

    return newItem, nil
}

func UpdateItem(request requests.UpdateItem, itemID string) (model.Item, error) {
	DB := config.DB
	
    item, err := FindByItemID(itemID)
	if err != nil {
		return item, err
	}

	var existingItem model.Item
	parseItemID, _ := uuid.FromString(itemID)
	err = DB.Where("name = ?", request.Name).Where("item_id != ?", parseItemID).First(&existingItem).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return existingItem, fmt.Errorf("Error find item by name")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return item, fmt.Errorf("Name item already exists")
	}
	

    item.Name = request.Name
    item.Description = request.Description

    if err := DB.Save(&item).Error; err != nil {
        return item, fmt.Errorf("Error update item")
    }

    return item, nil
}

func DeleteItem(itemID string) (model.Item, error) {
	DB := config.DB
	
	item, err := FindByItemID(itemID)
	if err != nil {
		return item, err
	}

	if err := DB.Delete(&item).Error; err!= nil {
        return item, fmt.Errorf("Error delete item")
    }

	return item, nil 
}