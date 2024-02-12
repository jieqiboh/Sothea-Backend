package models

import (
	"database/sql"
	"fmt"
)

type VisualAcuity struct {
	ID                     int32          `json:"id"`
	LEyeVision             int            `json:"lEyeVision" binding:"required"`
	REyeVision             int            `json:"rEyeVision" binding:"required"`
	AdditionalIntervention sql.NullString `json:"additionalIntervention"`
	//AdminID              uint   `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin                Admin
}

// TableName specifies the table name for the VisualAcuity model.
func (VisualAcuity) TableName() string {
	return "visualacuity"
}

// ToString generates a simple string representation of the VisualAcuity struct.
func (va VisualAcuity) String() string {
	additionalIntervention, _ := va.AdditionalIntervention.Value()
	return fmt.Sprintf("ID: %d\nLEyeVision: %d\nREyeVision: %d\nAdditionalIntervention: %s",
		va.ID, va.LEyeVision, va.REyeVision, additionalIntervention)
}
