package entities

import (
	"fmt"
)

type VisualAcuity struct {
	ID                     int32   `json:"id"`
	VID                    int32   `json:"vid" binding:"-"`
	LEyeVision             *int32  `json:"lEyeVision" binding:"required"`
	REyeVision             *int32  `json:"rEyeVision" binding:"required"`
	AdditionalIntervention *string `json:"additionalIntervention"`
	//AdminID              uint   `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin                Admin
}

// TableName specifies the table name for the VisualAcuity model.
func (VisualAcuity) TableName() string {
	return "visualacuity"
}

// ToString generates a simple string representation of the VisualAcuity struct.
func (va VisualAcuity) String() string {
	result := fmt.Sprintf("\nVISUAL ACUITY\n")
	result += fmt.Sprintf("ID: %d\n", va.ID)
	result += fmt.Sprintf("VID: %d\n", va.VID)
	result += fmt.Sprintf("LEyeVision: %d\n", *va.LEyeVision)
	result += fmt.Sprintf("REyeVision: %d\n", *va.REyeVision)
	result += fmt.Sprintf("Additional Intervention: %s\n", SafeDeref(va.AdditionalIntervention))
	return result
}
