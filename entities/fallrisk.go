package entities

import (
	"fmt"
)

type FallRisk struct {
	ID                 int32   `json:"id" binding:"-"`
	VID                int32   `json:"vid" binding:"-"`
	FallWorries        *string `json:"fallWorries" binding:"required"`
	FallHistory        *string `json:"fallHistory" binding:"required"`
	CognitiveStatus    *string `json:"cognitiveStatus" binding:"required"`
	ContinenceProblems *string `json:"continenceProblems" binding:"required"`
	SafetyAwareness    *string `json:"safetyAwareness" binding:"required"`
	Unsteadiness       *string `json:"unsteadiness" binding:"required"`
	FallRiskScore      *int32  `json:"fallRiskScore" binding:"required"`
}

// TableName specifies the table name for the FallRisk model.
func (FallRisk) TableName() string {
	return "fallrisk"
}

// ToString generates a simple string representation of the FallRisk struct.
func (fr FallRisk) String() string {
	result := fmt.Sprintf("\nFALL RISK\n")
	result += fmt.Sprintf("ID: %d\n", fr.ID)
	result += fmt.Sprintf("VID: %d\n", fr.VID)
	result += fmt.Sprintf("FallWorries: %s\n", *fr.FallWorries)
	result += fmt.Sprintf("FallHistory: %s\n", *fr.FallHistory)
	result += fmt.Sprintf("CognitiveStatus: %s\n", *fr.CognitiveStatus)
	result += fmt.Sprintf("ContinenceProblems: %s\n", *fr.ContinenceProblems)
	result += fmt.Sprintf("SafetyAwareness: %s\n", *fr.SafetyAwareness)
	result += fmt.Sprintf("Unsteadiness: %s\n", *fr.Unsteadiness)
	result += fmt.Sprintf("FallRiskScore: %d\n", *fr.FallRiskScore)
	return result
}
