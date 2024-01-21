package models

import "fmt"

type VitalStatistics struct {
	ID          int64   `json:"id"`
	Temperature float64 `json:"temperature" binding:"required"`
	SpO2        float64 `json:"spO2" binding:"required"`
	//AdminID     uint    `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin       Admin
}

// TableName specifies the table name for the VitalStatistics model.
func (VitalStatistics) TableName() string {
	return "vitalstatistics"
}

// ToString generates a simple string representation of the VitalStatistics struct.
func (vs VitalStatistics) ToString() string {
	return fmt.Sprintf("ID: %d, Temperature: %.2f, SpO2: %.2f",
		vs.ID, vs.Temperature, vs.SpO2)
}
