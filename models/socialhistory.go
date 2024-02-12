package models

import (
	"database/sql"
	"fmt"
)

type SocialHistory struct {
	ID                    int32          `json:"id"`
	PastSmokingHistory    bool           `json:"pastSmokingHistory" validate:"exists"`
	NumberOfYears         sql.NullInt32  `json:"numberOfYears"`
	CurrentSmokingHistory bool           `json:"currentSmokingHistory" validate:"exists"`
	CigarettesPerDay      sql.NullInt32  `json:"cigarettesPerDay"`
	AlcoholHistory        bool           `json:"alcoholHistory" validate:"exists"`
	HowRegular            sql.NullString `json:"howRegular"`
	//AdminID               uint    `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin                 Admin
}

// TableName specifies the table name for the SocialHistory model.
func (SocialHistory) TableName() string {
	return "socialhistory"
}

// ToString generates a simple string representation of the SocialHistory struct.
func (sh SocialHistory) String() string {
	// todo: handle errors
	numberOfYears, _ := sh.NumberOfYears.Value()
	cigarettesPerDay, _ := sh.CigarettesPerDay.Value()
	howRegular, _ := sh.HowRegular.Value()

	return fmt.Sprintf("ID: %d\nPastSmokingHistory: %t\nNumberOfYears: %v\nCurrentSmokingHistory: %t\nCigarettesPerDay: %v\n"+
		"AlcoholHistory: %t\nHowRegular: %v",
		sh.ID, sh.PastSmokingHistory, numberOfYears, sh.CurrentSmokingHistory, cigarettesPerDay,
		sh.AlcoholHistory, howRegular)
}
