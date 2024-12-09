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
func (p Physiotherapy) String() string {
	result := fmt.Sprintf("\nPHYSIOTHERAPY\n")
	result += fmt.Sprintf("ID: %d\n", p.ID)
	result += fmt.Sprintf("VID: %d\n", p.VID)
	result += fmt.Sprintf("PainStiffnessDay: %d\n", *p.PainStiffnessDay)
	result += fmt.Sprintf("PainStiffnessNight: %d\n", *p.PainStiffnessNight)
	result += fmt.Sprintf("SymptomsInterfereTasks: %s\n", *p.SymptomsInterfereTasks)
	result += fmt.Sprintf("SymptomsChange: %s\n", *p.SymptomsChange)
	result += fmt.Sprintf("SymptomsNeedHelp: %s\n", *p.SymptomsNeedHelp)
	result += fmt.Sprintf("TroubleSleepSymptoms: %s\n", *p.TroubleSleepSymptoms)
	result += fmt.Sprintf("HowMuchFatigue: %d\n", *p.HowMuchFatigue)
	result += fmt.Sprintf("AnxiousLowMood: %d\n", *p.AnxiousLowMood)
	result += fmt.Sprintf("MedicationManageSymptoms: %s\n", *p.MedicationManageSymptoms)
	return result
}
