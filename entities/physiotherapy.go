package entities

import (
	"fmt"
)

type Physiotherapy struct {
	ID                       int32   `json:"id" binding:"-"`
	VID                      int32   `json:"vid" binding:"-"`
	PainStiffnessDay         *int32  `json:"painStiffnessDay" binding:"required"`
	PainStiffnessNight       *int32  `json:"painStiffnessNight" binding:"required"`
	SymptomsInterfereTasks   *string `json:"symptomsInterfereTasks" binding:"required"`
	SymptomsChange           *string `json:"symptomsChange" binding:"required"`
	SymptomsNeedHelp         *string `json:"symptomsNeedHelp" binding:"required"`
	TroubleSleepSymptoms     *string `json:"troubleSleepSymptoms" binding:"required"`
	HowMuchFatigue           *int32  `json:"howMuchFatigue" binding:"required"`
	AnxiousLowMood           *int32  `json:"anxiousLowMood" binding:"required"`
	MedicationManageSymptoms *string `json:"medicationManageSymptoms" binding:"required"`
}

// TableName specifies the table name for the Physiotherapy model.
func (Physiotherapy) TableName() string {
	return "fallrisk"
}

// ToString generates a simple string representation of the Physiotherapy struct.
func (phy Physiotherapy) String() string {
	result := fmt.Sprintf("\nPHYSIOTHERAPY\n")
	result += fmt.Sprintf("ID: %d\n", phy.ID)
	result += fmt.Sprintf("VID: %d\n", phy.VID)
	result += fmt.Sprintf("PainStiffnessDay: %d\n", *phy.PainStiffnessDay)
	result += fmt.Sprintf("PainStiffnessNight: %d\n", *phy.PainStiffnessNight)
	result += fmt.Sprintf("SymptomsInterfereTasks: %s\n", *phy.SymptomsInterfereTasks)
	result += fmt.Sprintf("SymptomsChange: %s\n", *phy.SymptomsChange)
	result += fmt.Sprintf("SymptomsNeedHelp: %s\n", *phy.SymptomsNeedHelp)
	result += fmt.Sprintf("TroubleSleepSymptoms: %s\n", *phy.TroubleSleepSymptoms)
	result += fmt.Sprintf("HowMuchFatigue: %d\n", *phy.HowMuchFatigue)
	result += fmt.Sprintf("AnxiousLowMood: %d\n", *phy.AnxiousLowMood)
	result += fmt.Sprintf("MedicationManageSymptoms: %s\n", *phy.MedicationManageSymptoms)
	return result
}
