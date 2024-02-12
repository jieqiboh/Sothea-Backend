package models

import (
	"database/sql"
	"fmt"
)

type DoctorsConsultation struct {
	ID                int32          `json:"id" binding:"-"`
	Healthy           bool           `json:"healthy" validate:"exists"`
	Msk               bool           `json:"msk" validate:"exists"`
	Cvs               bool           `json:"cvs" validate:"exists"`
	Respi             bool           `json:"respi" validate:"exists"`
	Gu                bool           `json:"gu" validate:"exists"`
	Git               bool           `json:"git" validate:"exists"`
	Eye               bool           `json:"eye" validate:"exists"`
	Derm              bool           `json:"derm" validate:"exists"`
	Others            bool           `json:"others" validate:"exists"`
	ConsultationNotes sql.NullString `json:"consultationNotes"`
	Diagnosis         sql.NullString `json:"diagnosis"`
	Treatment         sql.NullString `json:"treatment"`
	ReferralNeeded    bool           `json:"referralNeeded" validate:"exists"`
	ReferralLoc       sql.NullString `json:"referralLoc"`
	Remarks           sql.NullString `json:"remarks"`
	// AdminID           uint    // Foreign key referencing Admin's ID
	// Admin             Admin
}

// TableName specifies the table name for the DoctorsConsultation model.
func (DoctorsConsultation) TableName() string {
	return "doctorsconsultation"
}

// ToString generates a simple string representation of the DoctorsConsultation struct.
func (d DoctorsConsultation) String() string {
	// todo: handle errors
	consultationNotes, _ := d.ConsultationNotes.Value()
	diagnosis, _ := d.Diagnosis.Value()
	treatment, _ := d.Treatment.Value()
	referralLoc, _ := d.ReferralLoc.Value()
	remarks, _ := d.Remarks.Value()

	return fmt.Sprintf("ID: %d\nHealthy: %t\nMsk: %t\nCvs: %t\nRespi: %t\nGu: %t\nGit: %t\nEye: %t\nDerm: %t\nOthers: "+
		"%t\nConsultationNotes: %s\nDiagnosis: %s\nTreatment: %s\nReferralNeeded: %t\nReferralLoc: %s\nRemarks: %s",
		d.ID, d.Healthy, d.Msk, d.Cvs, d.Respi, d.Gu, d.Git, d.Eye, d.Derm, d.Others,
		consultationNotes, diagnosis, treatment, d.ReferralNeeded, referralLoc, remarks)
}
