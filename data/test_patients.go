package data

import (
	"github.com/jieqiboh/sothea_backend/domain"
	"github.com/jieqiboh/sothea_backend/models"
	"time"
)

var admin1 = models.Admin{
	FamilyGroup:         models.PtrTo("S001"),
	RegDate:             models.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Name:                models.PtrTo("John Doe"),
	Dob:                 models.PtrTo(time.Date(1994, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Age:                 models.PtrTo(30),
	Gender:              models.PtrTo("M"),
	Village:             models.PtrTo("SO"),
	ContactNo:           models.PtrTo("12345678"),
	Pregnant:            models.PtrTo(false),
	LastMenstrualPeriod: models.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
	DrugAllergies:       models.PtrTo("panadol"),
	SentToID:            models.PtrTo(false),
}
var pastmedicalhistory1 = models.PastMedicalHistory{
	Tuberculosis:               models.PtrTo(true),
	Diabetes:                   models.PtrTo(false),
	Hypertension:               models.PtrTo(true),
	Hyperlipidemia:             models.PtrTo(false),
	ChronicJointPains:          models.PtrTo(false),
	ChronicMuscleAches:         models.PtrTo(true),
	SexuallyTransmittedDisease: models.PtrTo(true),
	SpecifiedSTDs:              models.PtrTo("TRICHOMONAS"),
	Others:                     nil,
}
var socialhistory1 = models.SocialHistory{
	PastSmokingHistory:    models.PtrTo(true),
	NumberOfYears:         models.PtrTo(int32(15)),
	CurrentSmokingHistory: models.PtrTo(false),
	CigarettesPerDay:      models.PtrTo(int32(5)),
	AlcoholHistory:        models.PtrTo(true),
	HowRegular:            models.PtrTo("A"),
}
var vitalstatistics1 = models.VitalStatistics{
	Temperature:              models.PtrTo(36.5),
	SpO2:                     models.PtrTo(98.0),
	SystolicBP1:              models.PtrTo(120.0),
	DiastolicBP1:             models.PtrTo(80.0),
	SystolicBP2:              models.PtrTo(122.0),
	DiastolicBP2:             models.PtrTo(78.0),
	AverageSystolicBP:        models.PtrTo(121.0),
	AverageDiastolicBP:       models.PtrTo(79.0),
	HR1:                      models.PtrTo(72.0),
	HR2:                      models.PtrTo(71.0),
	AverageHR:                models.PtrTo(71.5),
	RandomBloodGlucoseMmolL:  models.PtrTo(5.4),
	RandomBloodGlucoseMmolLp: models.PtrTo(5.3),
}
var heightandweight1 = models.HeightAndWeight{
	Height:      models.PtrTo(170.0),
	Weight:      models.PtrTo(70.0),
	BMI:         models.PtrTo(24.2),
	BMIAnalysis: models.PtrTo("normal weight"),
	PaedsHeight: models.PtrTo(90.0),
	PaedsWeight: models.PtrTo(80.0),
}
var visualacuity1 = models.VisualAcuity{
	LEyeVision:             models.PtrTo(int32(20)),
	REyeVision:             models.PtrTo(int32(20)),
	AdditionalIntervention: models.PtrTo("VISUAL FIELD TEST REQUIRED"),
}
var doctorsconsultation1 = models.DoctorsConsultation{
	Healthy:           models.PtrTo(true),
	Msk:               models.PtrTo(false),
	Cvs:               models.PtrTo(false),
	Respi:             models.PtrTo(true),
	Gu:                models.PtrTo(true),
	Git:               models.PtrTo(false),
	Eye:               models.PtrTo(true),
	Derm:              models.PtrTo(false),
	Others:            models.PtrTo(false),
	ConsultationNotes: models.PtrTo("CHEST PAIN, SHORTNESS OF BREATH, COUGH"),
	Diagnosis:         models.PtrTo("ACUTE BRONCHITIS"),
	Treatment:         models.PtrTo("REST, HYDRATION, COUGH SYRUP"),
	ReferralNeeded:    models.PtrTo(false),
	ReferralLoc:       models.PtrTo("REFERRED TO BOC"),
	Remarks:           models.PtrTo("MONITOR FOR RESOLUTION"),
}

// patient1 has all fields initialised and filled in
var Patient1 = domain.Patient{
	Admin:               &admin1,
	PastMedicalHistory:  &pastmedicalhistory1,
	SocialHistory:       &socialhistory1,
	VitalStatistics:     &vitalstatistics1,
	HeightAndWeight:     &heightandweight1,
	VisualAcuity:        &visualacuity1,
	DoctorsConsultation: &doctorsconsultation1,
}

// patient2 has all optional fields initialised to nil
var admin2 = models.Admin{
	FamilyGroup:         models.PtrTo("S002"),
	RegDate:             models.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Name:                models.PtrTo("Jane Smith"),
	Dob:                 models.PtrTo(time.Date(1994, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Age:                 models.PtrTo(3),
	Gender:              models.PtrTo("M"),
	Village:             models.PtrTo("SO"),
	ContactNo:           models.PtrTo("12345878"),
	Pregnant:            models.PtrTo(true),
	LastMenstrualPeriod: models.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
	DrugAllergies:       models.PtrTo("panadol"),
	SentToID:            models.PtrTo(true),
}
var pastmedicalhistory2 = models.PastMedicalHistory{
	Tuberculosis:               models.PtrTo(true),
	Diabetes:                   models.PtrTo(false),
	Hypertension:               models.PtrTo(true),
	Hyperlipidemia:             models.PtrTo(false),
	ChronicJointPains:          models.PtrTo(false),
	ChronicMuscleAches:         models.PtrTo(true),
	SexuallyTransmittedDisease: models.PtrTo(true),
	SpecifiedSTDs:              models.PtrTo("TRICHOMONAS"),
	Others:                     nil,
}
var socialhistory2 = models.SocialHistory{
	PastSmokingHistory:    models.PtrTo(true),
	NumberOfYears:         models.PtrTo(int32(5)),
	CurrentSmokingHistory: models.PtrTo(false),
	CigarettesPerDay:      models.PtrTo(int32(25)),
	AlcoholHistory:        models.PtrTo(true),
	HowRegular:            models.PtrTo("D"),
}
var vitalstatistics2 = models.VitalStatistics{
	Temperature:              models.PtrTo(36.5),
	SpO2:                     models.PtrTo(98.0),
	SystolicBP1:              models.PtrTo(10.0),
	DiastolicBP1:             models.PtrTo(80.0),
	SystolicBP2:              models.PtrTo(122.0),
	DiastolicBP2:             models.PtrTo(78.0),
	AverageSystolicBP:        models.PtrTo(11.0),
	AverageDiastolicBP:       models.PtrTo(79.0),
	HR1:                      models.PtrTo(72.0),
	HR2:                      models.PtrTo(7.0),
	AverageHR:                models.PtrTo(71.5),
	RandomBloodGlucoseMmolL:  models.PtrTo(5.4),
	RandomBloodGlucoseMmolLp: models.PtrTo(5.3),
}
var heightandweight2 = models.HeightAndWeight{
	Height:      models.PtrTo(170.0),
	Weight:      models.PtrTo(70.0),
	BMI:         models.PtrTo(24.2),
	BMIAnalysis: models.PtrTo("normal weight"),
	PaedsHeight: models.PtrTo(90.0),
	PaedsWeight: models.PtrTo(80.0),
}
var visualacuity2 = models.VisualAcuity{
	LEyeVision:             models.PtrTo(int32(20)),
	REyeVision:             models.PtrTo(int32(20)),
	AdditionalIntervention: models.PtrTo("VISUAL FIELD TEST REQUIRED"),
}
var doctorsconsultation2 = models.DoctorsConsultation{
	Healthy:           models.PtrTo(true),
	Msk:               models.PtrTo(false),
	Cvs:               models.PtrTo(false),
	Respi:             models.PtrTo(true),
	Gu:                models.PtrTo(true),
	Git:               models.PtrTo(false),
	Eye:               models.PtrTo(true),
	Derm:              models.PtrTo(false),
	Others:            models.PtrTo(false),
	ConsultationNotes: models.PtrTo("CHEST PAIN, SHORTNESS OF BREATH, COUGH"),
	Diagnosis:         models.PtrTo("ACUTE BRONCHITIS"),
	Treatment:         models.PtrTo("REST, HYDRATION, COUGH SYRUP"),
	ReferralNeeded:    models.PtrTo(true),
	ReferralLoc:       models.PtrTo("REFERRED TO BOC"),
	Remarks:           models.PtrTo("MONITOR FOR RESOLUTION"),
}
var Patient2 = domain.Patient{
	Admin:               &admin2,
	PastMedicalHistory:  &pastmedicalhistory2,
	SocialHistory:       &socialhistory2,
	VitalStatistics:     &vitalstatistics2,
	HeightAndWeight:     &heightandweight2,
	VisualAcuity:        &visualacuity2,
	DoctorsConsultation: &doctorsconsultation2,
}

// patient3 does only has admin filled in
var admin3 = models.Admin{
	FamilyGroup:         models.PtrTo("S002"),
	RegDate:             models.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Name:                models.PtrTo("Jane Smith"),
	Dob:                 models.PtrTo(time.Date(1994, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Age:                 models.PtrTo(3),
	Gender:              models.PtrTo("M"),
	Village:             models.PtrTo("SO"),
	ContactNo:           models.PtrTo("12345878"),
	Pregnant:            models.PtrTo(true),
	LastMenstrualPeriod: models.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
	DrugAllergies:       models.PtrTo("panadol"),
	SentToID:            models.PtrTo(true),
}
var Patient3 = domain.Patient{
	Admin:               &admin3,
	PastMedicalHistory:  nil,
	SocialHistory:       nil,
	VitalStatistics:     nil,
	HeightAndWeight:     nil,
	VisualAcuity:        nil,
	DoctorsConsultation: nil,
}

// patient4 is missing visualacuity and socialhistory categories
