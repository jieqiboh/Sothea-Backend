package models

import (
	"database/sql"
	"fmt"
)

type DoctorsConsultation struct {
	ID                int64          `json:"id" binding:"-"`
	Healthy           *bool          `json:"healthy" validate:"exists"`
	ConsultationNotes sql.NullString `json:"consultationnotes"`
	ReferralNeeded    *bool          `json:"referralneeded" validate:"exists"`
	//AdminID           uint    // Foreign key referencing Admin's ID
	//Admin             Admin
}

// TableName specifies the table name for the DoctorsConsultation model.
func (DoctorsConsultation) TableName() string {
	return "doctorsconsultation"
}

// ToString generates a simple string representation of the DoctorsConsultation struct.
func (d DoctorsConsultation) ToString() string {
	return fmt.Sprintf("ID: %d, Healthy: %t, ConsultationNotes: %s, ReferralNeeded: %t",
		d.ID, *d.Healthy, d.ConsultationNotes, *d.ReferralNeeded)
}
