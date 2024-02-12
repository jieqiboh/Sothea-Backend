package models

import (
	"database/sql"
	"fmt"
)

type PastMedicalHistory struct {
	ID                         int32          `json:"id" binding:"-"`
	Tuberculosis               bool           `json:"tuberculosis" validate:"exists"`
	Diabetes                   bool           `json:"diabetes" validate:"exists"`
	Hypertension               bool           `json:"hypertension" validate:"exists"`
	Hyperlipidemia             bool           `json:"hyperlipidemia" validate:"exists"`
	ChronicJointPains          bool           `json:"chronicJointPains" validate:"exists"`
	ChronicMuscleAches         bool           `json:"chronicMuscleAches" validate:"exists"`
	SexuallyTransmittedDisease bool           `json:"sexuallyTransmittedDisease" validate:"exists"`
	SpecifiedSTDs              sql.NullString `json:"specifiedSTDs"`
	Others                     sql.NullString `json:"others"`
	//AdminID                    uint `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin                      Admin
}

// TableName specifies the table name for the PastMedicalHistory model.
func (PastMedicalHistory) TableName() string {
	return "pastmedicalhistory"
}

// ToString generates a simple string representation of the PastMedicalHistory struct.
func (pmh PastMedicalHistory) String() string {
	// todo: handle errors
	specifiedSTDs, _ := pmh.SpecifiedSTDs.Value()
	others, _ := pmh.Others.Value()
	return fmt.Sprintf("ID: %d\nTuberculosis: %t\nDiabetes: %t\nHypertension: %t\nHyperlipidemia: %t\nChronicJointPains: %t\n"+
		"ChronicMuscleAches: %t\nSexuallyTransmittedDisease: %t\nSpecifiedSTDs: %s\nOthers: %s",
		pmh.ID, pmh.Tuberculosis, pmh.Diabetes, pmh.Hypertension, pmh.Hyperlipidemia, pmh.ChronicJointPains,
		pmh.ChronicMuscleAches, pmh.SexuallyTransmittedDisease, specifiedSTDs, others)
}
