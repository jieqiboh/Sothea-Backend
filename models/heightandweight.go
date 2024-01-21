package models

import "fmt"

type HeightAndWeight struct {
	ID     int64   `json:"id" binding:"-"`
	Height float64 `json:"height" binding:"required"`
	Weight float64 `json:"weight" binding:"required"`
	//AdminID uint    `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin   Admin
}

// TableName specifies the table name for the HeightAndWeight model.
func (HeightAndWeight) TableName() string {
	return "heightandweight"
}

// ToString generates a simple string representation of the HeightAndWeight struct.
func (hw HeightAndWeight) ToString() string {
	return fmt.Sprintf("ID: %d, Height: %.2f, Weight: %.2f",
		hw.ID, hw.Height, hw.Weight)
}
