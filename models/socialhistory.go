package models

import (
	"database/sql"
	"fmt"
)

type SocialHistory struct {
	ID                    int64         `json:"id"`
	PastSmokingHistory    *bool         `json:"pastsmokinghistory" validate:"exists"`
	NumberOfYears         sql.NullInt32 `json:"numberofyears"`
	CurrentSmokingHistory *bool         `json:"currentsmokinghistory" validate:"exists"`
	//AdminID               uint `gorm:"uniqueIndex;not null"` // Foreign key referencing Admin's ID
	//Admin                 Admin
}

// TableName specifies the table name for the SocialHistory model.
func (SocialHistory) TableName() string {
	return "socialhistory"
}

// ToString generates a simple string representation of the SocialHistory struct.
func (sh SocialHistory) ToString() string {
	return fmt.Sprintf("ID: %d, PastSmokingHistory: %t, NumberOfYears: %d, CurrentSmokingHistory: %t",
		sh.ID, *sh.PastSmokingHistory, sh.NumberOfYears, *sh.CurrentSmokingHistory)
}
