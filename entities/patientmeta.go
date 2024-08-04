package entities

import (
	"fmt"
	"time"
)

type PatientMeta struct {
	ID          int32       `json:"id" binding:"-"`
	VID         int32       `json:"vid" binding:"-"`
	FamilyGroup *string     `json:"familyGroup" binding:"required"`
	RegDate     *time.Time  `json:"regDate" binding:"required"`
	QueueNo     *string     `json:"queueNo" binding:"required"`
	Name        *string     `json:"name" binding:"required"`
	KhmerName   *string     `json:"khmerName" binding:"required"`
	Gender      *string     `json:"gender" binding:"required"`
	Village     *string     `json:"village" binding:"required"`
	ContactNo   *string     `json:"contactNo" binding:"required"`
	Visits      []time.Time `json:"visits"` // Slice of visit dates, 1-indexed according to the visit id
}

// TableName specifies the table name for the PatientMeta model.
func (PatientMeta) TableName() string {
	return "patientmeta"
}

// ToString generates a simple string representation of the PatientMeta struct.
func (pm PatientMeta) String() string {
	result := fmt.Sprintf("\nPATIENTMETA\n")
	result += fmt.Sprintf("ID: %d\n", pm.ID)
	result += fmt.Sprintf("VID: %d\n", pm.VID)
	result += fmt.Sprintf("FamilyGroup: %s\n", *pm.FamilyGroup)
	result += fmt.Sprintf("RegDate: %s\n", pm.RegDate.Format("2006-01-02"))
	result += fmt.Sprintf("QueueNo: %s\n", *pm.QueueNo)
	result += fmt.Sprintf("Name: %s\n", *pm.Name)
	result += fmt.Sprintf("KhmerName: %s\n", *pm.KhmerName)
	result += fmt.Sprintf("Gender: %s\n", *pm.Gender)
	result += fmt.Sprintf("Village: %s\n", *pm.Village)
	result += fmt.Sprintf("ContactNo: %s\n", *pm.ContactNo)
	result += fmt.Sprintf("Visits: %v\n", pm.Visits)
	return result
}
