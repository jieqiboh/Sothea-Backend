package entities

import (
	"fmt"
)

type SocialHistory struct {
	ID                    int32   `json:"id"`
	VID                   int32   `json:"vid" binding:"-"`
	PastSmokingHistory    *bool   `json:"pastSmokingHistory" binding:"required"`
	NumberOfYears         *int32  `json:"numberOfYears"`
	CurrentSmokingHistory *bool   `json:"currentSmokingHistory" binding:"required"`
	CigarettesPerDay      *int32  `json:"cigarettesPerDay"`
	AlcoholHistory        *bool   `json:"alcoholHistory" binding:"required"`
	HowRegular            *string `json:"howRegular"`
	//AdminID               uint    `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin                 Admin
}

// TableName specifies the table name for the SocialHistory model.
func (SocialHistory) TableName() string {
	return "socialhistory"
}

// ToString generates a simple string representation of the SocialHistory struct.
func (sh SocialHistory) String() string {
	result := fmt.Sprintf("\nSOCIAL HISTORY\n")
	result += fmt.Sprintf("ID: %d\n", sh.ID)
	result += fmt.Sprintf("VID: %d\n", sh.VID)
	result += fmt.Sprintf("Past Smoking History: %t\n", *sh.PastSmokingHistory)
	result += fmt.Sprintf("Number of Years: %d\n", SafeDeref(sh.NumberOfYears))
	result += fmt.Sprintf("Current Smoking History: %v\n", *sh.CurrentSmokingHistory)
	result += fmt.Sprintf("Cigarettes Per Day: %d\n", SafeDeref(sh.CigarettesPerDay))
	result += fmt.Sprintf("Alcohol History: %t\n", *sh.AlcoholHistory)
	result += fmt.Sprintf("How Regular: %v\n", SafeDeref(sh.HowRegular))
	return result
}
