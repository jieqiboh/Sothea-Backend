package models

import "fmt"

type VisualAcuity struct {
	ID         int64 `json:"id"`
	LEyeVision int   `json:"leyevision" binding:"required"`
	REyeVision int   `json:"reyevision" binding:"required"`
	//AdminID    uint `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin      Admin
}

// TableName specifies the table name for the VisualAcuity model.
func (VisualAcuity) TableName() string {
	return "visualacuity"
}

// ToString generates a simple string representation of the VisualAcuity struct.
func (va VisualAcuity) ToString() string {
	return fmt.Sprintf("ID: %d, LEyeVision: %d, REyeVision: %d",
		va.ID, va.LEyeVision, va.REyeVision)
}
