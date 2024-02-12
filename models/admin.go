package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Admin struct {
	ID                  int32          `json:"id" binding:"-"`
	FamilyGroup         string         `json:"familyGroup" binding:"required"`
	RegDate             time.Time      `json:"regDate" binding:"-"`
	Name                string         `json:"name" binding:"required"`
	Dob                 time.Time      `json:"dob" binding:"required"`
	Age                 int            `json:"age" binding:"required"`
	Gender              string         `json:"gender" binding:"required"`
	Village             string         `json:"village" binding:"required"`
	ContactNo           string         `json:"contactNo" binding:"required"`
	Pregnant            bool           `json:"pregnant" binding:"required"`
	LastMenstrualPeriod sql.NullTime   `json:"lastMenstrualPeriod"`
	DrugAllergies       sql.NullString `json:"drugAllergies"`
	SentToID            bool           `json:"sentToId" binding:"required"`
}

// TableName specifies the table name for the Admin model.
func (Admin) TableName() string {
	return "admin"
}

// ToString generates a simple string representation of the Admin struct.
func (a Admin) String() string {
	// todo: handle errors
	lastMenstrualPeriod, _ := a.LastMenstrualPeriod.Value()
	drugAllergies, _ := a.DrugAllergies.Value()

	return fmt.Sprintf("ID: %d\nFamilyGroup: %s\nRegDate: %s\nName: %s\nDob: %s\nAge: %d\nGender: %s\nVillage: %s\nContactNo: %s\nPregnant: %t\nLastMenstrualPeriod: %v\nDrugAllergies: %v\nSentToID: %t",
		a.ID, a.FamilyGroup, a.RegDate.Format("2006-01-02"), a.Name, a.Dob.Format("2006-01-02"), a.Age, a.Gender, a.Village, a.ContactNo, a.Pregnant, lastMenstrualPeriod, drugAllergies, a.SentToID)
}
