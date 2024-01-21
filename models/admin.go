package models

import (
	"fmt"
	"time"
)

type Admin struct {
	ID          int64     `json:"id" binding:"-"`
	FamilyGroup string    `json:"familygroup" binding:"required"`
	RegDate     time.Time `json:"regdate" binding:"-"`
	Name        string    `json:"name" binding:"required"`
	Age         int       `json:"age" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
}

// TableName specifies the table name for the Admin model.
func (a *Admin) TableName() string {
	return "admin"
}

// ToString generates a simple string representation of the Admin struct.
func (a Admin) ToString() string {
	return fmt.Sprintf("ID: %d, FamilyGroup: %s, RegDate: %s, Name: %s, Age: %d, Gender: %s",
		a.ID, a.FamilyGroup, a.RegDate.Format("2006-01-02"), a.Name, a.Age, a.Gender)
}
