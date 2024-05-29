package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jieqiboh/sothea_backend/controllers/middleware"
	"github.com/jieqiboh/sothea_backend/entities"
	"net/http"
	"strconv"
	"strings"
)

// PatientHandler represent the httphandler for patient
type PatientHandler struct {
	AUsecase entities.PatientUseCase
}

// NewPatientHandler will initialize the patients/ resources endpoint
func NewPatientHandler(e *gin.Engine, us entities.PatientUseCase) {
	handler := &PatientHandler{
		AUsecase: us,
	}

	// Protected routes
	authorized := e.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/patient/:id", handler.GetPatientByID)
		authorized.POST("/patient", handler.InsertPatient)
		authorized.DELETE("/patient/:id", handler.DeletePatientByID)
		authorized.PATCH("/patient/:id", handler.UpdatePatientByID)
		authorized.GET("/get-all-admin", handler.GetAllAdmin)
		authorized.GET("/search-patients/:search-name", handler.SearchPatients)
	}
}

func (p *PatientHandler) GetPatientByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))

	// Check if the id is not a number
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	id := int32(idP)
	ctx := c.Request.Context()

	// Get the patient by id
	patient, err := p.AUsecase.GetPatientByID(ctx, id)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (p *PatientHandler) InsertPatient(c *gin.Context) {
	var patient entities.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		// Use type assertion to check if err is of type validator.ValidationErrors
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Iterate over the validation errors
			for _, fieldErr := range validationErrs {
				c.JSON(http.StatusBadRequest, gin.H{"error": fieldErr.Error()})
				return // exit on first error
			}
		} else {
			// Handle other types of errors (e.g., JSON binding errors)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	ctx := c.Request.Context()
	id, err := p.AUsecase.InsertPatient(ctx, &patient)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "Inserted userid": id})
}

func (p *PatientHandler) DeletePatientByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	id := int32(idP)
	ctx := c.Request.Context()

	id, err = p.AUsecase.DeletePatientByID(ctx, id)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func (p *PatientHandler) UpdatePatientByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, entities.ErrPatientNotFound.Error())
	}

	id := int32(idP)
	ctx := c.Request.Context()

	var patient entities.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		// Use type assertion to check if err is of type validator.ValidationErrors
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Iterate over the validation errors
			for _, fieldErr := range validationErrs {
				c.JSON(http.StatusBadRequest, gin.H{"error": fieldErr.Error()})
				return // exit on first error
			}
		} else {
			// Handle other types of errors (e.g., JSON binding errors)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	id, err = p.AUsecase.UpdatePatientByID(ctx, id, &patient)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func (p *PatientHandler) GetAllAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	adminlist, err := p.AUsecase.GetAllAdmin(ctx)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, adminlist)
}

func (p *PatientHandler) SearchPatients(c *gin.Context) {
	ctx := c.Request.Context()

	patientName := c.Param("search-name")

	foundPatients, err := p.AUsecase.SearchPatients(ctx, patientName)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foundPatients)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case entities.ErrInternalServerError:
		return http.StatusInternalServerError
	case entities.ErrPatientNotFound:
		return http.StatusNotFound
	case entities.ErrMissingAdminCategory:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

type fieldError struct {
	err validator.FieldError
}

func (q fieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + q.err.Field() + "'")
	sb.WriteString(", condition: " + q.err.ActualTag())

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}
