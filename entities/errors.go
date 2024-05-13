package entities

import "errors"

var (
	// Should know which category has incorrect inputs
	// Should know which categories are missing
	ErrInternalServerError  = errors.New("Internal Server Error")
	ErrAuthenticationFailed = errors.New("Not Authenticated")
	ErrPatientNotFound      = errors.New("Patient Not Found")
	ErrMissingAdminCategory = errors.New("Missing Admin field")
	ErrInvalidInput         = errors.New("Invalid Parameters Given")
	ErrInvalidAdminFields   = errors.New("Invalid Admin Fields")
)
