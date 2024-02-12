package models

import "fmt"

type VitalStatistics struct {
	ID                       int32    `json:"id"`
	Temperature              *float64 `json:"temperature" binding:"required"`
	SpO2                     *float64 `json:"spO2" binding:"required"`
	SystolicBP1              *float64 `json:"systolicBP1" binding:"required"`
	DiastolicBP1             *float64 `json:"diastolicBP1" binding:"required"`
	SystolicBP2              *float64 `json:"systolicBP2" binding:"required"`
	DiastolicBP2             *float64 `json:"diastolicBP2" binding:"required"`
	AverageSystolicBP        *float64 `json:"averageSystolicBP" binding:"required"`
	AverageDiastolicBP       *float64 `json:"averageDiastolicBP" binding:"required"`
	HR1                      *float64 `json:"hr1" binding:"required"`
	HR2                      *float64 `json:"hr2" binding:"required"`
	AverageHR                *float64 `json:"averageHR" binding:"required"`
	RandomBloodGlucoseMmolL  *float64 `json:"randomBloodGlucoseMmolL" binding:"required"`
	RandomBloodGlucoseMmolLp *float64 `json:"randomBloodGlucoseMmolLp" binding:"required"`
	//AdminID                  uint    `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin                    Admin
}

// TableName specifies the table name for the VitalStatistics model.
func (VitalStatistics) TableName() string {
	return "vitalstatistics"
}

// ToString generates a simple string representation of the VitalStatistics struct.
func (vs VitalStatistics) String() string {
	result := fmt.Sprintf("\nVITAL STATISTICS\n")
	result += fmt.Sprintf("ID: %d\n", vs.ID)
	result += fmt.Sprintf("Temperature: %.1f\n", *vs.Temperature)
	result += fmt.Sprintf("SpO2: %.1f\n", *vs.SpO2)
	result += fmt.Sprintf("SystolicBP1: %.1f\n", *vs.SystolicBP1)
	result += fmt.Sprintf("DiastolicBP1: %.1f\n", *vs.DiastolicBP1)
	result += fmt.Sprintf("SystolicBP2: %.1f\n", *vs.SystolicBP2)
	result += fmt.Sprintf("DiastolicBP2: %.1f\n", *vs.DiastolicBP2)
	result += fmt.Sprintf("AverageSystolicBP: %.1f\n", *vs.AverageSystolicBP)
	result += fmt.Sprintf("AverageDiastolicBP: %.1f\n", *vs.AverageDiastolicBP)
	result += fmt.Sprintf("HR1: %.1f\n", *vs.HR1)
	result += fmt.Sprintf("HR2: %.1f\n", *vs.HR2)
	result += fmt.Sprintf("AverageHR: %.1f\n", *vs.AverageHR)
	result += fmt.Sprintf("RandomBloodGlucoseMmolL: %.1f\n", *vs.RandomBloodGlucoseMmolL)
	result += fmt.Sprintf("RandomBloodGlucoseMmolLp: %.1f\n", *vs.RandomBloodGlucoseMmolLp)
	return result
}
