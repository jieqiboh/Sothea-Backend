package entities

import (
	"fmt"
)

type Dental struct {
	ID               int32   `json:"id" binding:"-"`
	VID              int32   `json:"vid" binding:"-"`
	CleanTeethFreq   *int    `json:"cleanTeethFreq" binding:"required"`
	SugarConsumeFreq *int    `json:"sugarConsumeFreq" binding:"required"`
	PastYearDecay    *bool   `json:"pastYearDecay" binding:"required"`
	BrushTeethPain   *bool   `json:"brushTeethPain" binding:"required"`
	DrinkOtherWater  *bool   `json:"drinkOtherWater" binding:"required"`
	DentalNotes      *string `json:"dentalNotes"`
	ReferralNeeded   *bool   `json:"referralNeeded" binding:"required"`
	ReferralLoc      *string `json:"referralLoc"`
	Tooth11          *bool   `json:"tooth11" binding:"required"`
	Tooth12          *bool   `json:"tooth12" binding:"required"`
	Tooth13          *bool   `json:"tooth13" binding:"required"`
	Tooth14          *bool   `json:"tooth14" binding:"required"`
	Tooth15          *bool   `json:"tooth15" binding:"required"`
	Tooth16          *bool   `json:"tooth16" binding:"required"`
	Tooth17          *bool   `json:"tooth17" binding:"required"`
	Tooth18          *bool   `json:"tooth18" binding:"required"`
	Tooth21          *bool   `json:"tooth21" binding:"required"`
	Tooth22          *bool   `json:"tooth22" binding:"required"`
	Tooth23          *bool   `json:"tooth23" binding:"required"`
	Tooth24          *bool   `json:"tooth24" binding:"required"`
	Tooth25          *bool   `json:"tooth25" binding:"required"`
	Tooth26          *bool   `json:"tooth26" binding:"required"`
	Tooth27          *bool   `json:"tooth27" binding:"required"`
	Tooth28          *bool   `json:"tooth28" binding:"required"`
	Tooth31          *bool   `json:"tooth31" binding:"required"`
	Tooth32          *bool   `json:"tooth32" binding:"required"`
	Tooth33          *bool   `json:"tooth33" binding:"required"`
	Tooth34          *bool   `json:"tooth34" binding:"required"`
	Tooth35          *bool   `json:"tooth35" binding:"required"`
	Tooth36          *bool   `json:"tooth36" binding:"required"`
	Tooth37          *bool   `json:"tooth37" binding:"required"`
	Tooth38          *bool   `json:"tooth38" binding:"required"`
	Tooth41          *bool   `json:"tooth41" binding:"required"`
	Tooth42          *bool   `json:"tooth42" binding:"required"`
	Tooth43          *bool   `json:"tooth43" binding:"required"`
	Tooth44          *bool   `json:"tooth44" binding:"required"`
	Tooth45          *bool   `json:"tooth45" binding:"required"`
	Tooth46          *bool   `json:"tooth46" binding:"required"`
	Tooth47          *bool   `json:"tooth47" binding:"required"`
	Tooth48          *bool   `json:"tooth48" binding:"required"`
}

// TableName specifies the table name for the Dental model.
func (Dental) TableName() string {
	return "dental"
}

// ToString generates a simple string representation of the Dental struct.
func (fr Dental) String() string {
	result := fmt.Sprintf("\nDENTAL\n")
	result += fmt.Sprintf("ID: %d\n", fr.ID)
	result += fmt.Sprintf("VID: %d\n", fr.VID)
	result += fmt.Sprintf("CleanTeethFreq: %d\n", *fr.CleanTeethFreq)
	result += fmt.Sprintf("SugarConsumeFreq: %d\n", *fr.SugarConsumeFreq)
	result += fmt.Sprintf("PastYearDecay: %t\n", *fr.PastYearDecay)
	result += fmt.Sprintf("BrushTeethPain: %t\n", *fr.BrushTeethPain)
	result += fmt.Sprintf("DrinkOtherWater: %t\n", *fr.DrinkOtherWater)
	result += fmt.Sprintf("DentalNotes: %s\n", SafeDeref(fr.DentalNotes))
	result += fmt.Sprintf("ReferralNeeded: %t\n", *fr.ReferralNeeded)
	result += fmt.Sprintf("ReferralLoc: %s\n", SafeDeref(fr.ReferralLoc))
	result += fmt.Sprintf("Tooth11: %t\n", *fr.Tooth11)
	result += fmt.Sprintf("Tooth12: %t\n", *fr.Tooth12)
	result += fmt.Sprintf("Tooth13: %t\n", *fr.Tooth13)
	result += fmt.Sprintf("Tooth14: %t\n", *fr.Tooth14)
	result += fmt.Sprintf("Tooth15: %t\n", *fr.Tooth15)
	result += fmt.Sprintf("Tooth16: %t\n", *fr.Tooth16)
	result += fmt.Sprintf("Tooth17: %t\n", *fr.Tooth17)
	result += fmt.Sprintf("Tooth18: %t\n", *fr.Tooth18)
	result += fmt.Sprintf("Tooth21: %t\n", *fr.Tooth21)
	result += fmt.Sprintf("Tooth22: %t\n", *fr.Tooth22)
	result += fmt.Sprintf("Tooth23: %t\n", *fr.Tooth23)
	result += fmt.Sprintf("Tooth24: %t\n", *fr.Tooth24)
	result += fmt.Sprintf("Tooth25: %t\n", *fr.Tooth25)
	result += fmt.Sprintf("Tooth26: %t\n", *fr.Tooth26)
	result += fmt.Sprintf("Tooth27: %t\n", *fr.Tooth27)
	result += fmt.Sprintf("Tooth28: %t\n", *fr.Tooth28)
	result += fmt.Sprintf("Tooth31: %t\n", *fr.Tooth31)
	result += fmt.Sprintf("Tooth32: %t\n", *fr.Tooth32)
	result += fmt.Sprintf("Tooth33: %t\n", *fr.Tooth33)
	result += fmt.Sprintf("Tooth34: %t\n", *fr.Tooth34)
	result += fmt.Sprintf("Tooth35: %t\n", *fr.Tooth35)
	result += fmt.Sprintf("Tooth36: %t\n", *fr.Tooth36)
	result += fmt.Sprintf("Tooth37: %t\n", *fr.Tooth37)
	result += fmt.Sprintf("Tooth38: %t\n", *fr.Tooth38)
	result += fmt.Sprintf("Tooth41: %t\n", *fr.Tooth41)
	result += fmt.Sprintf("Tooth42: %t\n", *fr.Tooth42)
	result += fmt.Sprintf("Tooth43: %t\n", *fr.Tooth43)
	result += fmt.Sprintf("Tooth44: %t\n", *fr.Tooth44)
	result += fmt.Sprintf("Tooth45: %t\n", *fr.Tooth45)
	result += fmt.Sprintf("Tooth46: %t\n", *fr.Tooth46)
	result += fmt.Sprintf("Tooth47: %t\n", *fr.Tooth47)
	result += fmt.Sprintf("Tooth48: %t\n", *fr.Tooth48)
	return result
}
