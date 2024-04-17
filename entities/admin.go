package entities

import (
	"fmt"
	"time"
)

type Admin struct {
	ID                  int32      `json:"id" binding:"-"`
	FamilyGroup         *string    `json:"familyGroup" binding:"required"`
	RegDate             *time.Time `json:"regDate" binding:"required"`
	Name                *string    `json:"name" binding:"required"`
	Dob                 *time.Time `json:"dob" binding:"required"`
	Age                 *int       `json:"age" binding:"required"`
	Gender              *string    `json:"gender" binding:"required"`
	Village             *string    `json:"village" binding:"required"`
	ContactNo           *string    `json:"contactNo" binding:"required"`
	Pregnant            *bool      `json:"pregnant" binding:"required"`
	LastMenstrualPeriod *time.Time `json:"lastMenstrualPeriod"`
	DrugAllergies       *string    `json:"drugAllergies"`
	SentToID            *bool      `json:"sentToId" binding:"required"`
}

// TableName specifies the table name for the Admin model.
func (Admin) TableName() string {
	return "admin"
}

// ToString generates a simple string representation of the Admin struct.
func (a Admin) String() string {
	result := fmt.Sprintf("\nADMIN\n")
	result += fmt.Sprintf("ID: %d\n", a.ID)
	result += fmt.Sprintf("FamilyGroup: %s\n", *a.FamilyGroup)
	result += fmt.Sprintf("RegDate: %s\n", a.RegDate.Format("2006-01-02"))
	result += fmt.Sprintf("Name: %s\n", *a.Name)
	result += fmt.Sprintf("Dob: %s\n", a.Dob.Format("2006-01-02"))
	result += fmt.Sprintf("Age: %d\n", *a.Age)
	result += fmt.Sprintf("Gender: %s\n", *a.Gender)
	result += fmt.Sprintf("Village: %s\n", *a.Village)
	result += fmt.Sprintf("ContactNo: %s\n", *a.ContactNo)
	result += fmt.Sprintf("Pregnant: %t\n", *a.Pregnant)
	result += fmt.Sprintf("LastMenstrualPeriod: %v\n", SafeDeref(a.LastMenstrualPeriod))
	result += fmt.Sprintf("DrugAllergies: %v\n", SafeDeref(a.DrugAllergies))
	result += fmt.Sprintf("SentToID: %t\n", *a.SentToID)
	return result
}
