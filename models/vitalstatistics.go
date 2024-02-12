package models

import "fmt"

type VitalStatistics struct {
	ID                       int32   `json:"id"`
	Temperature              float64 `json:"temperature" binding:"required"`
	SpO2                     float64 `json:"spO2" binding:"required"`
	SystolicBP1              float64 `json:"systolicBP1" binding:"required"`
	DiastolicBP1             float64 `json:"diastolicBP1" binding:"required"`
	SystolicBP2              float64 `json:"systolicBP2" binding:"required"`
	DiastolicBP2             float64 `json:"diastolicBP2" binding:"required"`
	AverageSystolicBP        float64 `json:"averageSystolicBP" binding:"required"`
	AverageDiastolicBP       float64 `json:"averageDiastolicBP" binding:"required"`
	HR1                      float64 `json:"hr1" binding:"required"`
	HR2                      float64 `json:"hr2" binding:"required"`
	AverageHR                float64 `json:"averageHR" binding:"required"`
	RandomBloodGlucoseMmolL  float64 `json:"randomBloodGlucoseMmolL" binding:"required"`
	RandomBloodGlucoseMmolLp float64 `json:"randomBloodGlucoseMmolLp" binding:"required"`
	//AdminID                  uint    `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin                    Admin
}

// TableName specifies the table name for the VitalStatistics model.
func (VitalStatistics) TableName() string {
	return "vitalstatistics"
}

// ToString generates a simple string representation of the VitalStatistics struct.
func (vs VitalStatistics) String() string {
	return fmt.Sprintf("ID: %d\nTemperature: %.2f\nSpO2: %.2f\nSystolicBP1: %.2f\nDiastolicBP1: %.2f\nSystolicBP2: %.2f\n"+
		"DiastolicBP2: %.2f\nAverageSystolicBP: %.2f\nAverageDiastolicBP: %.2f\nHR1: %.2f\nHR2: %.2f\nAverageHR: %.2f\n"+
		"RandomBloodGlucoseMmolL: %.2f\nRandomBloodGlucoseMmolLp: %.2f",
		vs.ID, vs.Temperature, vs.SpO2, vs.SystolicBP1, vs.DiastolicBP1, vs.SystolicBP2, vs.DiastolicBP2, vs.AverageSystolicBP,
		vs.AverageDiastolicBP, vs.HR1, vs.HR2, vs.AverageHR, vs.RandomBloodGlucoseMmolL, vs.RandomBloodGlucoseMmolLp)
}
