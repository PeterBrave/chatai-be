package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type BasicTool struct {
	gorm.Model
	Name    string `json:"name"`
	Desc    string `json:"description"`
	Prompt  string `json:"prompt"`
}

// Create a new BasicTool record
func (b *BasicTool) Create() error {
	result := db.Create(&b)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return errors.New("failed to create BasicTool record")
	}
	return nil
}

// Retrieve a BasicTool record by its ID
func GetBasicToolByName(name string) (*BasicTool, error) {
	db = db.Debug()
	var b BasicTool
	result := db.Where("name = ?", name).First(&b)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("BasicTool record not found")
		}
		return nil, result.Error
	}
	return &b, nil
}

// Retrieve a list of all BasicTool records
func GetAllBasicTools() ([]BasicTool, error) {
	var tools []BasicTool
	result := db.Find(&tools)
	if result.Error != nil {
		return nil, result.Error
	}
	return tools, nil
}

// Update a BasicTool record
func (b *BasicTool) Update() error {
	result := db.Save(&b)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return errors.New("failed to update BasicTool record")
	}
	return nil
}

// Delete a BasicTool record by its ID
func DeleteBasicToolByID(id uint) error {
	result := db.Delete(&BasicTool{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return errors.New("failed to delete BasicTool record")
	}
	return nil
}
