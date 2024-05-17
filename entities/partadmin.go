package entities

import (
	"fmt"
	"time"
)

type PartAdmin struct {
	ID        int32      `json:"id" binding:"-"`
	Name      *string    `json:"name" binding:"required"`
	KhmerName *string    `json:"khmerName" binding:"required"`
	Dob       *time.Time `json:"dob" binding:"required"`
	Gender    *string    `json:"gender" binding:"required"`
	ContactNo *string    `json:"contactNo" binding:"required"`
}

// TableName specifies the table name for the PartAdmin model.
func (PartAdmin) TableName() string {
	return "partadmin"
}

// ToString generates a simple string representation of the PartAdmin struct.
func (pa PartAdmin) String() string {
	result := fmt.Sprintf("\nADMIN\n")
	result += fmt.Sprintf("ID: %d\n", pa.ID)
	result += fmt.Sprintf("Name: %s\n", *pa.Name)
	result += fmt.Sprintf("KhmerName: %s\n", *pa.KhmerName)
	result += fmt.Sprintf("Dob: %s\n", pa.Dob.Format("2006-01-02"))
	result += fmt.Sprintf("Gender: %s\n", *pa.Gender)
	result += fmt.Sprintf("ContactNo: %s\n", *pa.ContactNo)
	return result
}
