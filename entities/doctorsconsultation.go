package entities

import (
	"fmt"
)

type DoctorsConsultation struct {
	ID                int32   `json:"id" binding:"-"`
	VID               int32   `json:"vid" binding:"-"`
	Healthy           *bool   `json:"healthy" binding:"required"`
	Msk               *bool   `json:"msk" binding:"required"`
	Cvs               *bool   `json:"cvs" binding:"required"`
	Respi             *bool   `json:"respi" binding:"required"`
	Gu                *bool   `json:"gu" binding:"required"`
	Git               *bool   `json:"git" binding:"required"`
	Eye               *bool   `json:"eye" binding:"required"`
	Derm              *bool   `json:"derm" binding:"required"`
	Others            *string `json:"others" binding:"required"`
	ConsultationNotes *string `json:"consultationNotes"`
	Diagnosis         *string `json:"diagnosis"`
	Treatment         *string `json:"treatment"`
	ReferralNeeded    *bool   `json:"referralNeeded" binding:"required"`
	ReferralLoc       *string `json:"referralLoc"`
	Remarks           *string `json:"remarks"`
	// AdminID           uint    // Foreign key referencing Admin's ID
	// Admin             Admin
}

// TableName specifies the table name for the DoctorsConsultation model.
func (DoctorsConsultation) TableName() string {
	return "doctorsconsultation"
}

// ToString generates a simple string representation of the DoctorsConsultation struct.
func (dc DoctorsConsultation) String() string {
	result := fmt.Sprintf("\nDOCTOR'S CONSULTATION\n")
	result += fmt.Sprintf("ID: %d\n", dc.ID)
	result += fmt.Sprintf("VID: %d\n", dc.VID)
	result += fmt.Sprintf("Healthy: %t\n", *dc.Healthy)
	result += fmt.Sprintf("Msk: %t\n", *dc.Msk)
	result += fmt.Sprintf("Cvs: %t\n", *dc.Cvs)
	result += fmt.Sprintf("Respi: %t\n", *dc.Respi)
	result += fmt.Sprintf("Gu: %t\n", *dc.Gu)
	result += fmt.Sprintf("Git: %t\n", *dc.Git)
	result += fmt.Sprintf("Eye: %t\n", *dc.Eye)
	result += fmt.Sprintf("Derm: %t\n", *dc.Derm)
	result += fmt.Sprintf("Others: %s\n", *dc.Others)
	result += fmt.Sprintf("ConsultationNotes: %s\n", SafeDeref(dc.ConsultationNotes))
	result += fmt.Sprintf("Diagnosis: %s\n", SafeDeref(dc.Diagnosis))
	result += fmt.Sprintf("Treatment: %s\n", SafeDeref(dc.Treatment))
	result += fmt.Sprintf("ReferralNeeded: %t\n", *dc.ReferralNeeded)
	result += fmt.Sprintf("ReferralLoc: %s\n", SafeDeref(dc.ReferralLoc))
	result += fmt.Sprintf("Remarks: %s\n", SafeDeref(dc.Remarks))
	return result
}
