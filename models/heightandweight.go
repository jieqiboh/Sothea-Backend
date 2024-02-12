package models

import "fmt"

type HeightAndWeight struct {
	ID          int32   `json:"id" binding:"-"`
	Height      float64 `json:"height" binding:"required"`
	Weight      float64 `json:"weight" binding:"required"`
	BMI         float64 `json:"bmi" binding:"required"`
	BMIAnalysis string  `json:"bmiAnalysis" binding:"required"`
	PaedsHeight float64 `json:"paedsHeight" binding:"required"`
	PaedsWeight float64 `json:"paedsWeight" binding:"required"`
	//AdminID      uint    `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin        Admin
}

// TableName specifies the table name for the HeightAndWeight model.
func (HeightAndWeight) TableName() string {
	return "heightandweight"
}

// ToString generates a simple string representation of the HeightAndWeight struct.
func (hw HeightAndWeight) String() string {
	return fmt.Sprintf("ID: %d\nHeight: %.2f\nWeight: %.2f\nBMI: %.2f\nBMI Analysis: %s\nPaeds Height: %.2f\nPaeds Weight: %.2f",
		hw.ID, hw.Height, hw.Weight, hw.BMI, hw.BMIAnalysis, hw.PaedsHeight, hw.PaedsWeight)
}
