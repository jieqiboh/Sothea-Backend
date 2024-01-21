package models

import "fmt"

type PastMedicalHistory struct {
	ID                int64 `json:"id" binding:"-"`
	Tuberculosis      *bool `json:"tuberculosis" validate:"exists"`
	Diabetes          *bool `json:"diabetes" validate:"exists"`
	Hyperlipidemia    *bool `json:"hyperlipidemia" validate:"exists"`
	Hypertension      *bool `json:"hypertension" validate:"exists"`
	ChronicJointPains *bool `json:"chronicjointpains" validate:"exists"`
	//AdminID           uint `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin             Admin
}

// TableName specifies the table name for the PastMedicalHistory model.
func (PastMedicalHistory) TableName() string {
	return "pastmedicalhistory"
}

// ToString generates a simple string representation of the PastMedicalHistory struct.
func (pmh PastMedicalHistory) ToString() string {
	return fmt.Sprintf("ID: %d, Tuberculosis: %t, Diabetes: %t, Hypertension: %t, Hyperlipidemia: %t, ChronicJointPains: %t",
		pmh.ID, *pmh.Tuberculosis, *pmh.Diabetes, *pmh.Hypertension, *pmh.Hyperlipidemia, *pmh.ChronicJointPains)
}
