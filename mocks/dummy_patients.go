package mocks

import (
	"github.com/jieqiboh/sothea_backend/entities"
	"time"
)

// Valid Patient JSON
var ValidPatientJson = `{
  "admin": {
    "familyGroup": "S001",
    "regDate": "2024-01-10T00:00:00Z",
    "name": "Patient's Name Here",
    "khmerName": "តតតតតតត",
    "dob": "1994-01-10T00:00:00Z",
    "age": 30,
    "gender": "M",
    "village": "SO",
    "contactNo": "12345678",
    "pregnant": false,
    "lastMenstrualPeriod": null,
    "drugAllergies": "panadol",
    "sentToID": false,
    "photo": "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"
  },
  "pastMedicalHistory": {
    "tuberculosis": true,
    "diabetes": false,
    "hypertension": true,
    "hyperlipidemia": false,
    "chronicJointPains": false,
    "chronicMuscleAches": true,
    "sexuallyTransmittedDisease": true,
    "specifiedSTDs": "TRICHOMONAS",
    "others": null
  },
  "socialHistory": {
    "pastSmokingHistory": true,
    "numberOfYears": 15,
    "currentSmokingHistory": false,
    "cigarettesPerDay": null,
    "alcoholHistory": true,
    "howRegular": "A"
  },
  "vitalStatistics": {
    "temperature": 36.5,
    "spO2": 98,
    "systolicBP1": 120,
    "diastolicBP1": 80,
    "systolicBP2": 122,
    "diastolicBP2": 78,
    "averageSystolicBP": 121,
    "averageDiastolicBP": 79,
    "hr1": 72,
    "hr2": 71,
    "averageHR": 71.5,
    "randomBloodGlucoseMmolL": 5.4,
    "randomBloodGlucoseMmolLp": 5.3
  },
  "heightAndWeight": {
    "height": 170,
    "weight": 70,
    "bmi": 24.2,
    "bmiAnalysis": "normal weight",
    "paedsHeight": 90,
    "paedsWeight": 80
  },
  "visualAcuity": {
    "lEyeVision": 20,
    "rEyeVision": 20,
    "additionalIntervention": "VISUAL FIELD TEST REQUIRED"
  },
  "doctorsConsultation": {
    "healthy": true,
    "msk": false,
    "cvs": false,
    "respi": true,
    "gu": true,
    "git": false,
    "eye": true,
    "derm": false,
    "others": "TRICHOMONAS VAGINALIS",
    "consultationNotes": "CHEST PAIN, SHORTNESS OF BREATH, COUGH",
    "diagnosis": "ACUTE BRONCHITIS",
    "treatment": "REST, HYDRATION, COUGH SYRUP",
    "referralNeeded": false,
    "referralLoc": null,
    "remarks": "MONITOR FOR RESOLUTION"
  }
}`

var admin = entities.Admin{
	FamilyGroup:         entities.PtrTo("S001"),
	RegDate:             entities.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Name:                entities.PtrTo("Patient's Name Here"),
	KhmerName:           entities.PtrTo("តតតតតតត"),
	Dob:                 entities.PtrTo(time.Date(1994, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Age:                 entities.PtrTo(30),
	Gender:              entities.PtrTo("M"),
	Village:             entities.PtrTo("SO"),
	ContactNo:           entities.PtrTo("12345678"),
	Pregnant:            entities.PtrTo(false),
	LastMenstrualPeriod: nil,
	DrugAllergies:       entities.PtrTo("panadol"),
	SentToID:            entities.PtrTo(false),
	Photo:               entities.PtrTo("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"),
}
var pastmedicalhistory = entities.PastMedicalHistory{
	Tuberculosis:               entities.PtrTo(true),
	Diabetes:                   entities.PtrTo(false),
	Hypertension:               entities.PtrTo(true),
	Hyperlipidemia:             entities.PtrTo(false),
	ChronicJointPains:          entities.PtrTo(false),
	ChronicMuscleAches:         entities.PtrTo(true),
	SexuallyTransmittedDisease: entities.PtrTo(true),
	SpecifiedSTDs:              entities.PtrTo("TRICHOMONAS"),
	Others:                     nil,
}
var socialhistory = entities.SocialHistory{
	PastSmokingHistory:    entities.PtrTo(true),
	NumberOfYears:         entities.PtrTo(int32(15)),
	CurrentSmokingHistory: entities.PtrTo(false),
	CigarettesPerDay:      nil,
	AlcoholHistory:        entities.PtrTo(true),
	HowRegular:            entities.PtrTo("A"),
}
var vitalstatistics = entities.VitalStatistics{
	Temperature:              entities.PtrTo(36.5),
	SpO2:                     entities.PtrTo(98.0),
	SystolicBP1:              entities.PtrTo(120.0),
	DiastolicBP1:             entities.PtrTo(80.0),
	SystolicBP2:              entities.PtrTo(122.0),
	DiastolicBP2:             entities.PtrTo(78.0),
	AverageSystolicBP:        entities.PtrTo(121.0),
	AverageDiastolicBP:       entities.PtrTo(79.0),
	HR1:                      entities.PtrTo(72.0),
	HR2:                      entities.PtrTo(71.0),
	AverageHR:                entities.PtrTo(71.5),
	RandomBloodGlucoseMmolL:  entities.PtrTo(5.4),
	RandomBloodGlucoseMmolLp: entities.PtrTo(5.3),
}
var heightandweight = entities.HeightAndWeight{
	Height:      entities.PtrTo(170.0),
	Weight:      entities.PtrTo(70.0),
	BMI:         entities.PtrTo(24.2),
	BMIAnalysis: entities.PtrTo("normal weight"),
	PaedsHeight: entities.PtrTo(90.0),
	PaedsWeight: entities.PtrTo(80.0),
}
var visualacuity = entities.VisualAcuity{
	LEyeVision:             entities.PtrTo(int32(20)),
	REyeVision:             entities.PtrTo(int32(20)),
	AdditionalIntervention: entities.PtrTo("VISUAL FIELD TEST REQUIRED"),
}
var doctorsconsultation = entities.DoctorsConsultation{
	Healthy:           entities.PtrTo(true),
	Msk:               entities.PtrTo(false),
	Cvs:               entities.PtrTo(false),
	Respi:             entities.PtrTo(true),
	Gu:                entities.PtrTo(true),
	Git:               entities.PtrTo(false),
	Eye:               entities.PtrTo(true),
	Derm:              entities.PtrTo(false),
	Others:            entities.PtrTo("TRICHOMONAS VAGINALIS"),
	ConsultationNotes: entities.PtrTo("CHEST PAIN, SHORTNESS OF BREATH, COUGH"),
	Diagnosis:         entities.PtrTo("ACUTE BRONCHITIS"),
	Treatment:         entities.PtrTo("REST, HYDRATION, COUGH SYRUP"),
	ReferralNeeded:    entities.PtrTo(false),
	ReferralLoc:       nil,
	Remarks:           entities.PtrTo("MONITOR FOR RESOLUTION"),
}
var ValidPatient = entities.Patient{
	Admin:               &admin,
	PastMedicalHistory:  &pastmedicalhistory,
	SocialHistory:       &socialhistory,
	VitalStatistics:     &vitalstatistics,
	HeightAndWeight:     &heightandweight,
	VisualAcuity:        &visualacuity,
	DoctorsConsultation: &doctorsconsultation,
}

// Missing Admin Field
var MissingAdminPatientJson = `{
  "pastMedicalHistory": {
    "tuberculosis": true,
    "diabetes": false,
    "hypertension": true,
    "hyperlipidemia": false,
    "chronicJointPains": false,
    "chronicMuscleAches": true,
    "sexuallyTransmittedDisease": true,
    "specifiedSTDs": "TRICHOMONAS",
    "others": null
  },
  "socialHistory": {
    "pastSmokingHistory": true,
    "numberOfYears": 15,
    "currentSmokingHistory": false,
    "cigarettesPerDay": null,
    "alcoholHistory": true,
    "howRegular": "A"
  },
  "vitalStatistics": {
    "temperature": 36.5,
    "spO2": 98,
    "systolicBP1": 120,
    "diastolicBP1": 80,
    "systolicBP2": 122,
    "diastolicBP2": 78,
    "averageSystolicBP": 121,
    "averageDiastolicBP": 79,
    "hr1": 72,
    "hr2": 71,
    "averageHR": 71.5,
    "randomBloodGlucoseMmolL": 5.4,
    "randomBloodGlucoseMmolLp": 5.3
  },
  "heightAndWeight": {
    "height": 170,
    "weight": 70,
    "bmi": 24.2,
    "bmiAnalysis": "normal weight",
    "paedsHeight": 90,
    "paedsWeight": 80
  },
  "visualAcuity": {
    "lEyeVision": 20,
    "rEyeVision": 20,
    "additionalIntervention": "VISUAL FIELD TEST REQUIRED"
  },
  "doctorsConsultation": {
    "healthy": true,
    "msk": false,
    "cvs": false,
    "respi": true,
    "gu": true,
    "git": false,
    "eye": true,
    "derm": false,
    "others": "TRICHOMONAS VAGINALIS",
    "consultationNotes": "CHEST PAIN, SHORTNESS OF BREATH, COUGH",
    "diagnosis": "ACUTE BRONCHITIS",
    "treatment": "REST, HYDRATION, COUGH SYRUP",
    "referralNeeded": false,
    "referralLoc": null,
    "remarks": "MONITOR FOR RESOLUTION"
  }
}`

var MissingAdminPatient = entities.Patient{
	PastMedicalHistory:  &pastmedicalhistory,
	SocialHistory:       &socialhistory,
	VitalStatistics:     &vitalstatistics,
	HeightAndWeight:     &heightandweight,
	VisualAcuity:        &visualacuity,
	DoctorsConsultation: &doctorsconsultation,
}

// Invalid Parameters
var InvalidParametersPatientJson = `{
  "admin": {
    "regDate": "2024-01-10T00:00:00Z",
    "name": "Patient's Name Here",
    "khmerName": "តតតតតតត",
    "dob": "1994-01-10T00:00:00Z",
    "age": 30,
    "gender": "M",
    "village": "SO",
    "contactNo": "12345678",
    "pregnant": false,
    "lastMenstrualPeriod": null,
    "drugAllergies": "panadol",
    "sentToID": false,
    "photo": "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"
  },
  "pastMedicalHistory": {
    "tuberculosis": "invalid data type here",
    "diabetes": false,
    "hypertension": true,
    "hyperlipidemia": false,
    "chronicJointPains": false,
    "chronicMuscleAches": true,
    "sexuallyTransmittedDisease": true,
    "specifiedSTDs": "TRICHOMONAS",
    "others": null
  },
  "socialHistory": {
    "pastSmokingHistory": true,
    "numberOfYears": 15,
    "currentSmokingHistory": false,
    "cigarettesPerDay": null,
    "alcoholHistory": true,
    "howRegular": "A"
  },
  "vitalStatistics": {
    "temperature": 36.5,
    "spO2": 98,
    "systolicBP1": 120,
    "diastolicBP1": 80,
    "systolicBP2": 122,
    "diastolicBP2": 78,
    "averageSystolicBP": 121,
    "averageDiastolicBP": 79,
    "hr1": 72,
    "hr2": 71,
    "averageHR": 71.5,
    "randomBloodGlucoseMmolL": 5.4,
    "randomBloodGlucoseMmolLp": 5.3
  },
  "heightAndWeight": {
    "height": 170,
    "weight": 70,
    "bmi": 24.2,
    "bmiAnalysis": "normal weight",
    "paedsHeight": 90,
    "paedsWeight": 80
  },
  "visualAcuity": {
    "lEyeVision": 20,
    "rEyeVision": 20,
    "additionalIntervention": "VISUAL FIELD TEST REQUIRED"
  },
  "doctorsConsultation": {
    "healthy": true,
    "msk": false,
    "cvs": false,
    "respi": true,
    "gu": true,
    "git": false,
    "eye": true,
    "derm": false,
    "others": "TRICHOMONAS VAGINALIS",
    "consultationNotes": "CHEST PAIN, SHORTNESS OF BREATH, COUGH",
    "diagnosis": "ACUTE BRONCHITIS",
    "treatment": "REST, HYDRATION, COUGH SYRUP",
    "referralNeeded": false,
    "referralLoc": null,
    "remarks": "MONITOR FOR RESOLUTION"
  }
}`

// JSON Marshalling Error
var JSONMarshallingErrorPatientJson = `{
  "admin": {
    "familyGroup": false,
    "regDate": "2024-01-10T00:00:00Z",
    "name": "Patient's Name Here",
    "khmerName": "តតតតតតត",
    "dob": "1994-01-10T00:00:00Z",
    "age": 30,
    "gender": "M",
    "village": "SO",
    "contactNo": "12345678",
    "pregnant": false,
    "lastMenstrualPeriod": null,
    "drugAllergies": "panadol",
    "sentToID": false,
    "photo": "iVBORw0KGgoAAAANSUhEUgAAAgAAAAIACAYAAAD0eNT6AAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAOxAAADsQBlSsOGwAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAACAASURBVHic7d13tCZVme/x7+k+nbuhiQLS0E2SIKKASFRAVEZxnKsy6OKOeXQMGMcxcc33mq8ZR1RUFHEUlAEUxyxRycEm04GcQ0N30/HcP/Z7ru8cTp9Utd+nwvez1m+dHtew6tn1Vu3alXaBJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJEmSJGki+qILkFTYJsC2wPadv9sAmwNbdv5uAswGZgFTgbk8cd9fBazo+vsI8FAnDwB3A3d0chuwCHgsY5skZeYAQKqPJwNPB/YEdgF27fzdLKCWAeBW4HpgYefvtcB1wIMB9UgaJwcAUjVtDBwIHADsTzrwbxFa0djdDVwCnAdcAFwKrA6tSNITOACQqmEj4FDgiM7fPYBJgfWUaSVwMWlAcD5wEbAstCJJkgI9FfgQ6Sx5DemyehuyFvg98HbScwuSJDVaH+mS/heAm4k/EFcllwLHkwZEkiQ1xlOB/016cj76YFv13AR8ivRwoyRJtTMHeCPpnnf0QbWuubSzDmeOc91LktRzewPfIb0jH30AbUruB74I7D6O30GSpOwmAy8DziX+YNn0nAcc01nnkiSFmAG8DVhC/IGxbbkFeCveHpAk9dBs4L3AXcQfCNuee4F/I01zLElSFtOB9wD3EX/gM/899wDvJl2VkSSpFP3AP5M+ghN9oDMj53bgdTRnFkVJUpDnAtcQf2Az48u1wNHD/J6SJI1oJ+BM4g9kpljOAnZAkqRRTAc+BjxO/MHLlJMVwIc7v60kSU9wGHAD8Qcskyc3As9BkqSOjUmz960n/iBl8mY98HXSVM2SpBZ7LrCU+AOT6W2WAocjSWqdGcBX8Ky/zVkHfA6YiiSpFXYDrib+AGSqkctJ24QkqcFehV/qM0/MCtIEQpKkhpkOfJf4A42pdr4BTEOS1AjzgIuJP7iYeuQSYFskSbV2COlDMdEHFVOv3Ak8C0lSLR0DrCT+YGLqmcdJz4xIjTQ5ugApgz7g46TX/KYE16L66gf+gTQY+FNwLZKkUUwBfkj82aNpVr6JJ0ySVFkzgV8Qf7Awzcx/krYxqRH6oguQSjKX9OnXg6MLUaOdBxwFLIsuRCrKAYCaYC7wW2Cf6ELUChcBLwQeji5EKsIBgOpuDvBrYP/oQipgJbAIWAws6eQe4H7gAeBe/nbmuoo0+90k0tcQpwCzSRMmzQA2J70LP6+TbTvZCefOhzR98PNJ61WqJQcAqrNZwDmkd/3b5iHgQtKB6OpObiF94CanKcDuwNOHZG7m5VbR5aSvSXolQJJ6aAbpsn/0g2G9yv3AqcA/kw7AVRq89wF7Ae8hDciWE7++epWLSFehJEk9MI10oInu/HPnGtJ8BvtTr1fQpgGHAZ8i3ZKIXo+580d8O0CSsptCeh0rutPPleuBj5HO8pugjzSA+RJpet3o9ZsrZ5ImDpIkZfIt4jv7srMCOJnmP8swGTgc+BGwmvj1XnZOLG9VSZK6vZv4Tr7MLOm0aZMS11FdbAN8kvRmQvTvUGY+UuZKkiSlyVfWEt/Bl5FLgFfgJWNIrx2+lvQWQ/TvUlb+Z6lrSJJabHfSq1bRHXvRXA0cTbWe4K+KPtK6uYH436loVuK8FJJU2JakS+XRnXqRXAu8FA/8Y9EPvB5YSvzvViR3kiZOkiRNwHTgAuI784nmAeDt+EniiZhGej7iEeJ/x4nm4k47JEnj9G3iO/GJZB1wArBZ+aukdbYBfkL8bzrRnFD+KpGkZnsF8Z33RHITaRIclesw4Drif9+J5FUZ1ockNdJuwKPEd9zjyRrSrH1e8s1nBvBp0hWW6N97PHkU2DXD+pCkRplO/V4Juxmf+u6lw4Dbif/dx5PL8SuKkjSiLxDfWY8nJ+HHYCJsBvyc+N9/PPlsljUhSQ1wKPW5vLsSeE2OlaBxeTNpKuXo7WEsWYfPh0jSE2xEfd73XwzsnWUtaCKeSX0+NLQEmJ1lLUhSTX2F+M55LPkTvt5XRduS7rNHbx9jydcyrQNJqp19qMc8/z8hPYmuapoF/Iz47WS0rAOenWkdSFJtTAYuJb5THi2fxKl862AS8Bnit5fRcgO+Miqp5d5OfGc8UtaTpqRVvbyP+G1ntByfrfWSVHFbAQ8R3xGPdPA/Llvrldt7SL9h9Ha0oawAFmRrvSRVWJXneF+Hr/k1wXFUexBwZr6mS1I1vYD4zndDWQ/8S76mq8feRLUHAc/L13RJqpY+4AriO94N5d/yNV1BPkT8drWhXEV6GFaSGq/KX/r7ZMZ2K9YJxG9fG8prM7Zbkiqhn/QKVHSHO1xOwVf9mmwycAbx29lwuZ30ISxJaqzXE9/ZDpdz8b3sNpgBnEf89jZc3pqx3ZIUahrVnO9/MU7v2yZbALcRv90NzVJgSsZ2S1KY44jvZIdmJbBvzkarkvYHVhG//Q3N63M2WpIizATuIr6DHZrXZGyzqu1dxG9/Q3MTvhEgqWGqePZ/UtYWq+r6gNOI3w6H5ticjZakXpoMLCK+Y+3OzcCcnI1WLWxEOuuO3h67s5D0USNJqr1/JL5T7c4a0j1gCeAQ0tTP0dtld16WtcWS1CN/Ib5D7c7H8zZXNfRV4rfL7vwpb3MlKb9nEt+Zdud6nHBFTzQLuIX47bM7u2VtsSRldhLxHelg1gEH522uauxwqvXRoM/lba4k5bMJsJz4jnQwJ+RtrhqgSgPWe3F2Skk1VaVX/x7A2f40uq2AR4nfXgdzTN7mSlIelxHfgQ7muMxtVXP8L+K318H8NnNbJal0exDfeQ7mWpxjXWM3i/R1vujtdoD0TMJOeZurtnKyCeXy6ugCuhxPevdfGovlwIeji+joA14bXYQkjVUf1fna2mWdeqTxmAxcRfz2OwBcl7mtklSaA4nvNAfzosxtVXMdTfz2OxjnBJBUC58nvsMcAC7J3VA12mTSNyOit+MB4AOZ2ypJpajKjGq+QqWi3kL8djwAXJy7oZJU1F7Ed5YDwBKgP29T1QIzSRPyRG/P64FtM7dVLeNbACrbC6ML6PgKsDa6CNXeCuDr0UWQHmR9SXQRkjSSc4k/W1pBmoZYKsPmwCrit2snBZJUWZuQ3reP7ihPzt1Qtc7pxG/Xq4HZuRuq9vAWgMp0GNW47/6t6ALUON+PLoA0m+X+0UWoORwAqEzPiS4AuAE4L7oINc45wH3RReDnrFUiBwAq06HRBQA/ji5AjbQGODW6COCQ6AIkaahNgXXE3yfdPXdD1Vp7E799Lwem5m6oJI3HUcR3jtdkb6XargozAz4reyvVCt4CUFmeGV0A8LPoAtR4v4wuAG8DqCQOAFSW/aILID2oJeX0q+gCgIOiC5CkbvcRe1n0ftLHW6ScZpAmmorc1u/K3kq1glcAVIbtSLOlRfo16SFEKaeVwB+Ca9gK2DK4BjWAAwCVYY/oAojvlNUeVbjVtGd0Aao/BwAqQxUGAOdHF6DWuCC6AOBp0QWo/hwAqAzRA4AHgeuDa1B7/JV0KyCS812oMAcAKsOuwcu/iPRwlNQLa4Arg2vYOXj5agAHACrDjsHLvzR4+Wqfi4OXv1Pw8tUADgBU1Gxgi+AanAFQvXZJ8PK3AWYF16CacwCgonaILoB0T1bqpcuCl98HbB9cg2rOAYCKmh+8/JWk+dmlXlpE/LwTzgWgQhwAqKhtgpd/C/EdsdpnNXB7cA1PCl6+as4BgIraKnj5S4KXr/ZaFLx8BwAqxAGAioruhBYHL1/tFT0A8BaACnEAoKKiO6GlwctXe0UPAKIH36o5BwAqam7w8u8OXr7a69bg5UcPvlVzDgBU1MbBy78/ePlqr4eDl+8AQIU4AFBRGwUv/4Hg5au9Hgle/qbBy1fNOQBQUdEDAK8AKEr0AGBK8PJVcw4AVNT04OUvD16+2mtZ8PInBy9fNecAQEVFn4WsDl6+2iv6CoADABXiAEBFRXdCDgAUZUXw8qMH36o5BwAqKroTcgCgKP3By48efKvmHACoqIHg5fcFL1/tFX0Ajl6+as4BgIqKPgOfEbx8tVf0FYDo5avmHACoqFXBy58VvHy1V/QZuP23CnEDUlFrgpcf/Rqi2iv6DHxl8PJVcw4AVFT0FYCZwctXe80OXn70WwiqOQcAKsoBgNoqei5+BwAqxAGAioqeDGWz4OWrvaIHAM6CqUIcAKio6Ln4tw1evtprq+Dl+yEsFeIAQEVFd0Lzgpev9npS8PKj9z3VnAMAFRV9BcABgKJsEbx8BwAqxAGAioruhLwFoCg7BC8/evCtmnMAoKLuC16+VwAUZbfg5d8RvHzVnAMAFbU0ePk7AlODa1D7zAK2C64het9TzTkAUFGLg5c/lfgzMbXPbsR/iOq24OWr5hwAqKglxH8R8OnBy1f77BpdAA4AVJADABW1Arg3uAYHAOq1PYKXfzfwWHANqjkHACpD9L1IBwDqtQOCl39D8PLVAA4AVIbozmgv4u/Hqj2mAvsF1xC9z6kBHACoDFcHL38T4GnBNag99gVmBNdwY/Dy1QAOAFSG6AEAwBHRBag1DokuALgqugDVnwMAleGa6AJwAKDeqcIA4MroAiRp0L2k1wGjshyYlr2VarupwMPEbuu3Z2+lWsErACpL9CXJmcCBwTWo+Q4FNg6u4Yrg5ashHACoLBdFFwC8ILoANd4/RBcAXBxdgCR1O5LYy6IDpFkJfR1QuUwC7iR+Oz88d0MlaTw2AtYS3znun7uhaq0DiN++1wKzczdU7eAtAJVlGbAwugjgFdEFqLGqcPn/apwCWCVxAKAyXRhdAHAMMDm6CDVOP3BsdBHAH6ILUHM4AFCZfh9dALAV8JzoItQ4LwKeHF0E8JvoAiRpOHOBNcTfJ/1R7oaqdc4mfrteBczK3VBJmqjziO8oVwPb5G6oWmMe1XjA9Q+5G6p28RaAyvar6AKAKcCbo4tQY7yeajxX8ovoAiRpJHsTf6Y0QJqaeHrmtqr5ppGm3o3engeAXTK3VZIK6aM6HeZrM7dVzfcm4rfjAeC63A2VpDJ8kfgOc4D0zrS3uTRR/cAi4rfjAeAzmdsqSaWowoxpg3FiIE3Uq4nffgezT+a2SlIp+kjz8kd3mgPAjaQzOWk8JgF/JX77HQBuztxWtZSXR5XDAHBadBEdO5PO5KTxOAbYI7qIjlOiC5Ck8diX+DOnwSwlPc0tjcUMqnMFawDYLWtrJSmDK4nvPAfzrsxtVXN8mPjtdTB/ztxWScribcR3oINZRjXmcle1bUv62l709jqYN+RtriTlsTGwnPhOdDD/kbe5aoAfE7+dDuZRYE7e5kpSPt8nviPtzpF5m6saezawnvhtdDDfzttcScrrYOI70u7cTHrIS+o2i7RtRG+f3XlG1hZLUg9cTHxn2p1P522uaugE4rfL7vwxa2slqUeOIb5D7c464LlZW6w6OYJqXfofAF6StcWS1CP9wGLiO9Xu3A5slrPRqoW5wG3Eb4/duZlqfH5YkkrxTuI71qH5edYWqw5OIX47HJo3ZW2xJPXYHOAB4jvXoXlzzkar0t5O/PY3NLfjrJWSGuiDxHewQ7MCeGbORquSngOsJn77G5p35my0JEWZDdxLfCc7NHeSZoBTO8wD7iF+uxuae0ivI0pSI/0r8R3tcLkCO982mE71XksdzHEZ2y1J4WaSzrijO9vh8jP8RHaTTaJaU/12ZzHe+5fUAm8hvsPdUD6Tsd2K9Q3it68N5dUZ2y1JldEP/JX4TndD+VC+pivIJ4nfrjaUK/G9f0ktcjjxHe9IeX++pqvHqvRZ6uFyRL6mS1I1/Sfxne9I8aGs+nsd1Zvmtzun5Wu6JFXXzsAq4jvhDWU9zspWZ++k2gf/lcCCbK2XpIqr8r3ZwUGAzwTUT9W3qwHg+Gytl6QamAYsJL4zHi3fIT28qGrrA75I/PYyWq7H1/4kiUNIn+iN7pRHyxnAjEzrQMVNp7rv+XdnHXBQpnUgSbVT5Xe0u3MBsEWmdaCJmwdcQvz2MZZ8PdM6kKRa2pjqfZd9Q7kNODDPatAEPIdqzu0/XG7GKacl6QmeDawlvpMeS9YA7yPdc1acN1LNr/oNl3WkbVySNIzPEt9RjyenARtlWRMayVzgFOJ///Hkc1nWhCQ1xFTgMuI76/HkJtKDjOqN51Of20WDuZS0bUuSRrAbsJz4Tns8WQ98E5iTYX0omQF8mnq8MdKdR4GnZFgfktRIrya+455IFgPPy7A+2u7ZpAfoon/fieSVGdaHJDXa14nvvCeS9cBJwFblr5LWmQecSrWn9B0pJ5S/SiSp+aYA5xLfiU80j5EuWXtbYPxmkN6yeJT433GiuRDv+0vShG0D3EV8Z14kt5O+Sjep5HXTRJNIl8yXEP+7FcmdpG1XklTAQaQvp0V36kVzFXA0MLnc1dMI/cCrgOuI/52KZiVOFCVJpTma+j39vaEsAt4BzCx1DdXTVNKB/0bif5cysh44ttQ1JEni/cR38GXmXuAjwOZlrqSamEdq+x3E/w5l5oNlriRJ0t+cSHwnX3YeB04H/p704GNTTSa18SzqM+XzeHJieatKkjRUP3Am8Z19rtwLfAl4RlkrLNgkYH/gM6SHIaPXb66cjs92SFJ2U4FziO/0c2cRaXbBo6nXq4STgYOBL9Psg/5gfg9ML2XNSZJGNQs4n/jOv1dZAfwCeBvp6kB/8VVYmqmkp97fR7q8/zDx66tXuYh6Dc6k/89PmarONgZ+B+wTXUiAlcAVpI/MXEz6gNIi0qdxc5oF7Er6XsMewAHAfqSJe9rmPOBFpMmKpNpxAKC62wz4L9o5CBhqHemS+6KuLCWdkS/r5JGubDzkv59C+szuFp1sDWzZyY6kA/922G9Auuz/96SPVkmSgswmdcjRl4NNO3IO7bziIUmVNJN0JSD64GCanbPxgT9JqpzpNPsVQROb02j2PA2SVGv9pFfnog8Wplk5gWq9eSFJ2oB30JxvB5i4rCW93ihJqpGXkd6fjz6ImHpmGXAUkqRaOpD0bfbog4mpV24mzXUgSaqxLfA1QTP2nE+a+0CS1ABTgK8Sf3Ax1c164GukqY0lSQ3zT8BjxB9sTLVyH2lmP0lSgy0ALiT+oGOqkd8DT0aS1Ar9wEdJr3lFH4BMTNaQtoHJSJJa5zDgVuIPRqa3uQV4FpKkVtsI+DpOHNSGrAW+BMxBkqSOA4GFxB+kTJ5cBuyHJEnDmAZ8DHic+AOWKScPAW8BJiFJ0ijmAScTf/AyxXIWsC2SJI3Tc4GriT+QmfHlUuDwYX5PSZLGbDLwL8AdxB/YzMhZSPoIVN+wv6QkSRMwFXgjflyoilnS+W18p1+SlM0s0jfi7yf+wNf23E66OjNlxF9MkqQSzQKOAxYRfyBsW64hnfHPGPVXkiQpk0nAi4ELiD8wNjnrgN901rX3+CVJlXIQ8D1gOfEHzKbkYeALwA5j/xkkSYqxMWnymcuJP4DWNReR7u/PHue6lySpEvYGPgssJv6gWvUsBj4J7DKhNS1JUgX1kb5A93/xC4TduQn4NLDvxFetpPHyQRopRh+wF3BkJwcB/aEV9c6jwJ+A35Ie6rs2thypnRwASNWwMXAEcChwMLAnzZnUZi3wF9IB/7fAnzv/m6RADgCkatoI2J90ZeAA0tWCLUMrGru7gEu6ciHprF9ShTgAkOrjSaQrA0/r/N0FWABsHVTPXaSH9q7t5K+d3BVUj6RxcAAg1d900kBgPrA96UrBZsME0q2GSaSpcwdfr1sDPNb59+PAsk4eAR4C7gHuJX0M6R7SvPuLgJXZWiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiQF6IsuQGqpGcACYHNgs87fLTr/7s5UYA7QT9pf53b++6nArN6WvEHLgdWdfz8MDABrgUc7//sDwP2dvw8A93X9+35gMbCytyVLcgAg5TEJ2KGT+aSD/fyubBVTVmXdDSzpZHHXv2/p/N/rY8qSmssBgFTcxsCewO7AHsA+wF7A7MiiGmQ1cDNwGbAQuBb4C3BvZFFS3TkAkMZnOrAvcBBwIPAMYF5oRe11G3AFcAFwIXAp8HhoRVKNOACQRrYl8CzSWf1BwMGkQYCqZy1wFWlAcD7wJ7xKIG2QAwDpv+sH9geOAo4A9sb9pM4WAWcDZwHnAatiy5Gqw45NSg/kPZ900H8+6Z6+mmcF6VbB2cAZwNLYcqRYDgDUVrsDxwAvIz24p/b5K3Aa8BPguuBapJ5zAKA2mQ+8BDiadD9fGnQt8FPgx8D1wbVIPeEAQE23HXAs8I/A04NrUT1cQboqcArpTQNJUk1MJj3A9xNgDWlmOmPGm3XAb0hXjKYgSaqsecD7gFuJP3iYZuUu4NPAjkiSKmEy8HLgd6TpYqMPFKbZWQ/8lrTNTUaS1HOzgTeSHtiKPiiYdmYR6YqTr41KUg88Cfgo6Uty0QcAYwaAR4AvA9siSSrdHsB3STO5RXf4xgyXx4GTSHNMSJIKWgB8kzTPe3QHb8xYso70BspTkCSN2/akA7+v8Zm6ZnAgsBOSpFHNI91PfZz4DtyYMrIaOBnYAUnSE2wCfBEP/Ka5eRz4AjAXSRKTgFcB9xDfQRvTizwAvIP06WlJaqXDgauI75CNich1wJFIUovsRHo4KroDNqYKOQunGFYAp7JUL00DPgycCjwtuBapKnYB3kTqjy8ivT0gSY1xILCQ+LMtY6qcG4FDkXpgUnQBaryZpK+onYszpEmj2Rn4PWkOjDnBtUjShL0QWEr8WZUxdcwdwP9AkmpkLnAK8R2oMU3ID/CLg8qgL7oANc4BwA9x1jOpTLcC/0S6lSaVwmcAVJZ+0md6z8ODv1S27YA/kKbJnhpcixrCKwAqwwLSWf+B0YVILXAJcCxwU3QhqjevAKioNwDX4MFf6pVnApcDrwmuQzXnFQBN1AzgG8CrowupueWkNyVuBW4D7gYe7OShrn+v6vz/ru78dytJH5epgumk7QHS5elZnf9tE2DTTgb/vTXpi4/zSJ98ntXrYhvmO8DbqM62oBpxAKCJmA/8DHhGcB11sQq4gTT3+7WdvzeSDvgPBtZVBZuSBgO7kOaJ2B3YDXgK3useq0uBl5EGkdKYOQDQeL2A9IrfZtGFVNRK4ArgYtK92suAW4C1kUXVUD9pfvx9SZe8n0kacM4Y6T9qsfuBVwK/jS5EUvP0AR8iHcii34uuUu4jfdjozcDT8ROvOU0hDQLeAvyUdNCL/v2rlLXA+/HETlKJZgNnEN/BVSGPAWcD7wL2ws420iTSgOA9wC9Iz0hEbx9VyOn4bIWkEmxNuscY3alF5j7gZOBo0mBI1TQdOIL0rvwdxG83kbkK2LbY6pTUZnuSHiyK7swishT4DLA/vi5bR5NIr6Z+lvZuw0uAPQquR0ktdATwMPGdWC/zMOlM/8Wkb7OrGSYBB5OuDNxH/HbWyywD/q74KpTUFq8nvW8e3Xn1IuuAX5Jeo5pexspTpU0HXg6cQ/rto7e/XmQ1ThokaQw+QXyH1YvcA3yaNI2x2mkH0m2etlwV+Ggpa01S4/QBXyW+k8qdC0nvS08rZ7WpAaaR5tb/M/HbZ+58Ed9ckdRlMvBt4junnDmfdG9fGsnBwFnAeuK32Vw5GZ9xkUTqCL5PfKeUI6tJk/TsU9raUlvsSTpQNvVZmFNJEytJaqlpNHOCnzXASaRvFkhFLAC+RzNnwPw53gqTWmkm8F/Ed0JlZh3wY9KHZKQy7Ua6mtS0WwO/IvUFklpiGs07+J8FPK3MlSQN4xmk10ajt/eyBwFeCZBaYDLpTCa60ykr1wMvLHUNSaN7LnAN8dt/WTkDP2QlNdok0qd8ozubMvIg8D78Zrzi9ANvpDnzCPwU3w6QGqkP+A7xnUzRrAP+Hdis3NUjTdjmwIk04/mAE3GeAKlxPk9851I0NwGHlb1ipJIcBCwkfj8pmq+UvWIkxan79L6rSNOY+qCSqm4a8HHSNhu93xTJR0teL5ICvIH4zqRILgJ2L32tSHk9FfgL8fvPRLMePyAk1dph1PdMZA3pYz3OVqa66ic9qFrX2QRXA88rfa1Iym534CHiO5GJZBFpTnapCfYDbiR+v5pIHsH5NaRa2Qa4lfjOYyL5FjC7/FUihZpDfd/CWQJsVfoakVS6mdTz3uMK4LUZ1odUJccCjxG/v403l+HAXKq0Pur5cZ/rgD0yrA+pip5KmsEyer8bb07HOQKkyjqe+E5ivPk5sHGOlSFV2BzgP4jf/8ab9+dYGZKKeQFplrzoDmKsWQu8O8uakOqhD3gv9dtvj8ixMiRNzPbUaz7yR4GXZFkTUv0cSXraPnq/HGseABZkWROSxmU6cCnxncJYczuwd5Y1IdXXnqSn7aP3z7HmCmBGjhUhaexOIr4zGGv+jK8TSRuyNfV6g+fkPKtB0ljUaZrfs0mvKErasJnAL4jfX8ea12RZC5JGtCOwjPgOYCw5Baf0lcaqn/pc2XsM2CXPapA0nH7S5fTonX8s+RowKc9qkBqrD/gC8fvvWHIJDvClnvk/xO/0Y8nxuVaA1BIfIX4/Hks+kWsFSPqbg0nv4kbv8KPlg7lWgNQy/0r8/jxa1gGHZmq/JNKMeUuI39lHy4cytV9qq/cSv1+PlluBTXKtAKntTiF+Jx8tH8jWeqnd6jDV9w+ytV5qsaOI37lHi2f+Ul4fJn4/Hy0vzNZ6qYXmkC6vRe/YI+Uz2Vovqdvnid/fR8pSUp8lqQRfJ36nHik/xM+ESr3SB3yX+P1+pHw5W+ulFtmfan8t7GzSvASSemcK8Evi9/8Nbup5uwAACYNJREFUZR1wULbWSy0wFVhI/M68ofwZmJWt9ZJGMhO4gPh+YEO5DpiWrfVSw32M+J14Q7kZ2Cxf0yWNwRbAIuL7gw3lI/maLjXXjsDjxO/Aw2UZ8NR8TZc0DrsBDxPfLwyXVcDO+ZouNdOZxO+8w2Ud6ZVESdVxJNWdIfRnGdstNc7hxO+0G8q7M7Zb0sRVebbA52Vst9QY/cA1xO+ww+X7GdstqbgfEN9PDJergMkZ2y01wluJ31mHyzWkp44lVdcM4Eri+4vh8qaM7ZZqbxPgPuJ31KF5lPSgkaTq24X0oG50vzE09wJzM7ZbqrUvEr+TDpdX5Gy0pNIdS3y/MVw+n7PRUl1tRzVf+/tazkZLyubfie8/hmYlsG3ORkt19C3id86hWUi6pyipfqZTzQeKT8jZaKlu5pMmzIjeMbuzGtg3Y5sl5bc3aV+O7k+G9i0LcjZaqpPvE79TDs0HsrZYUq98mPj+ZGi+nbXFUk3sDKwhfofszoX4zq7UFP2kD3dF9yvdWQs8JWejpTo4lfidsTvLgZ2ytlhSr+0CrCC+f+nOD7K2WKq4PUhz60fviN35t6wtlhTlA8T3L91ZC+yatcVShX2P+J2wO1cDU3I2WFKYfqo3S+C3srZYqqhtqNaT/+uAA7K2WFK0Z1Gtq46PA1tnbbFUQZ8ifufrzlfzNldSRVRtgqBP5G2uVC2zgQeJ3/EGcyewUdYWS6qKucA9xPc7g7kfmJW1xVKFvIP4na47r8vbXEkV80bi+53uvDVvc6VqmAzcTPwON5gr8Z1/qW0mA1cR3/8MZhH2Q2qBlxO/s3Xn8LzNlVRRRxDf/3TnpXmbK8X7HfE72mDOyNxWSdV2FvH90GB+nbmtUqidgPXE72gDpOmHnYpTarddSRPyRPdHA6TXE3fI21x1mxRdQMu8HuiLLqLjB8AN0UVICnU98MPoIjom4QPJaqh+0ut20aPsAdLnOB1pSwLYkep8MvguUl8pNcpLid+5BvPNzG2VVC/fJr5fGsyLM7dV6rlziN+xBkjTD8/P21RJNbM91Zma/MzMbZV6ajuq86DNCZnbKqmeTiS+fxog9ZVPztxWqWc+SPxONbhj7Zi5rZLqaWeq86Gg92Zuq9QzVfkE509zN1RSrf2c+H5qALg0d0OlXngK8TvTYPzcr6SRHEx8PzWYnTK3tfWcByC/Y6IL6DgfuCi6CEmVdj7w5+giOo6OLkAq6hriR9IDOM+2pLE5mvj+aoB061SqrV2J34kGgNvwS1uSxqYfuJ34fmsA2C1zW1vNWwB5vTK6gI7vkJ7ulaTRrAW+G11Eh7cBVFsLiR9BrwXm5W6opEaZTzVeCbw6czulLOYRv/MMkD73KUnj9Svi+68B0kRqysBbAPm8KLqAjhOjC5BUS9+KLqDj+dEFSON1BvEjZ7+sJWmipgD3Et+PnZ67oVKZpgLLiN9xvpS7oZIa7QTi+7FHSIMRqRaeS/xOM4Az/0kq5hDi+7EB4Nm5G9pGPgOQx5HRBQC3Up0ZvSTV0/nA0ugiqEaf2jgOAPL4u+gCgB+RRs6SNFEDVOMjYg4AVAtPIv5y2QDwjNwNldQK+xDfn60HNs3d0LbxCkD5DoougHT5/4roIiQ1wmXE3wboAw4MrqFxHACUrwob6ZnRBUhqlHOiC6AaJ1eN4gCgfFXYSH8RXYCkRqlCn3JwdAHSSGYAq4i9V/YYMD13QyW1ygxgObF92+PAtNwNbROvAJRrX9IkQJF+Q9pRJKksK4E/BNcwDdg7uIZGcQBQrirc//9VdAGSGqkKzwF4G6BEDgDKVYUBwB+jC5DUSH+MLoBq9LHSsG4j9h7ZXfmbKKml+oB7iO3jluRuZJt4BaA8mwLbBtfwp+DlS2quAeC84Bq2A+YG19AYDgDK87ToAoBzowuQ1GjRJxl9wJ7BNTSGA4DyVGEAEL1zSmq2KvQxDgBK4gCgPE8NXv7DwLXBNUhqtr8Cy4JrcABQEgcA5dkrePmXk+7RSVIu64n/zkgVrrY2ggOAckwC9giu4bLg5Utqh+i+Zk/SswAqyAFAORYAs4JriB6VS2qHy4OXPwfYPriGRnAAUI6dogsgflQuqR2iBwAAO0YX0AQOAMoxP3j5y4Cbg2uQ1A43AI8G1zA/ePmN4ACgHPODl38t6eEcScptPfFvHC0IXn4jOAAox/zg5d8QvHxJ7XJT8PLnBy+/ERwAlCN6NBq9M0pqlxuDlx/d5zaCA4ByzA9efvTOKKldoq86zg9efiM4AChuBrBlcA0OACT1UnSfszUwPbiG2nMAUNx8YielGMBbAJJ660ZiZx7tw7kACuuPLqABnhS8/D5geXANktRrWxJ/K6LWvAJQ3GbRBUhSC9n3FuQAoDg3QknqPfveghwAFLd5dAGS1EL2vQU5ACjOUagk9Z59b0EOAIpzI5Sk3rPvLcgBQHFehpKk3rPvLcgBQHGbRhcgSS3kFYCCHAAUNzO6AElqIfveghwAFDctugBJaqGp0QXUnQOA4twIJan3PPkqyAFAcQ4AJKn37HsLcgBQnKNQSeo9+96CHAAU5yhUknrPvrcgBwDFuRFKUu95BaAgBwDFOQCQpN5zAFCQA4Di1kUXIEkttDa6gLpzAFDcsugCJKmFHokuoO4cABS3OLoASWqhRdEF1J0DgOKujC5AklroqugC6s4BQHG/jy5Aklrod9EF1F1fdAENMBu4G5gVXYgktcRyYCvgsehC6swrAMU9Bvw4ughJapEf4cG/MK8AlGNnYCEwJboQSWq41cBu+BBgYZOjC2iIB4E5wEHRhUhSw30OOC26iCbwCkB5pgN/APaPLkSSGuoi4DBgVXQhTeAAoFxbARcD86ILkaSGuRPYD7gjupCm8CHAct0NHAXcHl2IJDXIbcCRePAvlQOA8l0N7A2cG12IJDXARaQz/2uiC2kaHwLMYwVwKrAe2Be/GChJ47Ua+CzwWpz3PwsHAPmsA/4InESaJGh3HAhI0miWA98DXkl62t8vrmbiQ4C9Mxt4EekJ1qcDC4C5OCiQ1F6rgYdJH1W7gvQm1S9xkh9JkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJkiRJUoX9P3Uj/DMSzt+lAAAAAElFTkSuQmCC"
  },
  "pastMedicalHistory": {
    "tuberculosis": true,
    "diabetes": false,
    "hypertension": true,
    "hyperlipidemia": false,
    "chronicJointPains": false,
    "chronicMuscleAches": true,
    "sexuallyTransmittedDisease": true,
    "specifiedSTDs": "TRICHOMONAS",
    "others": null
  },
  "socialHistory": {
    "pastSmokingHistory": true,
    "numberOfYears": 15,
    "currentSmokingHistory": false,
    "cigarettesPerDay": null,
    "alcoholHistory": true,
    "howRegular": "A"
  },
  "vitalStatistics": {
    "temperature": 36.5,
    "spO2": 98,
    "systolicBP1": 120,
    "diastolicBP1": 80,
    "systolicBP2": 122,
    "diastolicBP2": 78,
    "averageSystolicBP": 121,
    "averageDiastolicBP": 79,
    "hr1": 72,
    "hr2": 71,
    "averageHR": 71.5,
    "randomBloodGlucoseMmolL": 5.4,
    "randomBloodGlucoseMmolLp": 5.3
  },
  "heightAndWeight": {
    "height": 170,
    "weight": 70,
    "bmi": 24.2,
    "bmiAnalysis": "normal weight",
    "paedsHeight": 90,
    "paedsWeight": 80
  },
  "visualAcuity": {
    "lEyeVision": 20,
    "rEyeVision": 20,
    "additionalIntervention": "VISUAL FIELD TEST REQUIRED"
  },
  "doctorsConsultation": {
    "healthy": true,
    "msk": false,
    "cvs": false,
    "respi": true,
    "gu": true,
    "git": false,
    "eye": true,
    "derm": false,
    "others": "TRICHOMONAS VAGINALIS",
    "consultationNotes": "CHEST PAIN, SHORTNESS OF BREATH, COUGH",
    "diagnosis": "ACUTE BRONCHITIS",
    "treatment": "REST, HYDRATION, COUGH SYRUP",
    "referralNeeded": false,
    "referralLoc": null,
    "remarks": "MONITOR FOR RESOLUTION"
  }
}`

// Admin Array
var AdminArray = []entities.PartAdmin{
	{
		ID:        1,
		Name:      entities.PtrTo("John Doe"),
		KhmerName: entities.PtrTo("១២៣៤ ៥៦៧៨៩០ឥឲ"),
		Dob:       entities.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
		Gender:    entities.PtrTo("M"),
		ContactNo: entities.PtrTo("12345678"),
	},
	{
		ID:        3,
		Name:      entities.PtrTo("Bob Smith"),
		KhmerName: entities.PtrTo("១២៣៤ ៥៦៧៨៩០ឥឲ"),
		Dob:       entities.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
		Gender:    entities.PtrTo("M"),
		ContactNo: entities.PtrTo("99999999"),
	},
	{
		ID:        4,
		Name:      entities.PtrTo("Bob Johnson"),
		KhmerName: entities.PtrTo("១២៣៤ ៥៦៧៨៩០ឥឲ"),
		Dob:       entities.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
		Gender:    entities.PtrTo("M"),
		ContactNo: entities.PtrTo("11111111"),
	},
	{
		ID:        5,
		Name:      entities.PtrTo("Alice Brown"),
		KhmerName: entities.PtrTo("១២៣៤ ៥៦៧៨៩០ឥឲ"),
		Dob:       entities.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
		Gender:    entities.PtrTo("F"),
		ContactNo: entities.PtrTo("17283948"),
	},
	{
		ID:        6,
		Name:      entities.PtrTo("Charlie Davis"),
		KhmerName: entities.PtrTo("១២៣៤ ៥៦៧៨៩០ឥឲ"),
		Dob:       entities.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
		Gender:    entities.PtrTo("M"),
		ContactNo: entities.PtrTo("09876543"),
	},
	{
		ID:        2,
		Name:      entities.PtrTo("New Patient's Name Here"),
		KhmerName: entities.PtrTo("តតតតតតត"),
		Dob:       entities.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
		Gender:    entities.PtrTo("M"),
		ContactNo: entities.PtrTo("12345678"),
	},
}
