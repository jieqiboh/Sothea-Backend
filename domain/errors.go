package domain

import "errors"

var (
	// Should know which category has incorrect inputs
	// Should know which categories are missing
	ErrInternalServerError  = errors.New("Internal Server Error")
	ErrAuthenticationFailed = errors.New("Not Authenticated")
	ErrNotFound             = errors.New("Patient Not Found")
	ErrMissingAdminInput    = errors.New("Missing Admin field")
	ErrInvalidInput         = errors.New("Invalid Parameters Given")
)

/**
 * List of possible errors
 * Errors that happen commonly include:
 */
