package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jieqiboh/sothea_backend/controllers/middleware"
	"github.com/jieqiboh/sothea_backend/entities"
	"github.com/jieqiboh/sothea_backend/util"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// PatientHandler represent the httphandler for patient
type PatientHandler struct {
	Usecase entities.PatientUseCase
}

// NewPatientHandler will initialize the patients/ resources endpoint
func NewPatientHandler(e *gin.Engine, us entities.PatientUseCase, secretKey []byte) {
	handler := &PatientHandler{
		Usecase: us,
	}

	// Protected routes
	authorized := e.Group("/")
	authorized.Use(middleware.AuthRequired(secretKey))
	{
		authorized.GET("/patient/:id/:vid", handler.GetPatientVisit)
		authorized.POST("/patient", handler.CreatePatient)
		authorized.POST("/patient/:id", handler.CreatePatientVisit)
		authorized.DELETE("/patient/:id/:vid", handler.DeletePatientVisit)
		authorized.PATCH("/patient/:id/:vid", handler.UpdatePatientVisit)
		authorized.GET("/patient-meta/:id", handler.GetPatientMeta)
		authorized.GET("/all-patient-visit-meta/:date", handler.GetAllPatientVisitMeta)
		authorized.GET("/export-db", handler.ExportDatabaseToCSV)
	}
}

func (p *PatientHandler) GetPatientVisit(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vidP, err := strconv.Atoi(c.Param("vid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := int32(idP)
	vid := int32(vidP)
	ctx := c.Request.Context()

	// Get the patient by id
	patient, err := p.Usecase.GetPatientVisit(ctx, id, vid)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (p *PatientHandler) CreatePatient(c *gin.Context) {
	var patientAdmin entities.Admin

	if err := c.ShouldBindJSON(&patientAdmin); err != nil {
		// Use type assertion to check if err is of type validator.ValidationErrors
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Get the first Validation Error
			fieldErr := validationErrs[0]
			c.JSON(http.StatusBadRequest, gin.H{"error": fieldErr.Error()})
			return // exit on first error
		} else if err.Error() == "EOF" {
			// Handle Empty Request Body Errors
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request Body is empty!"})
			return
		} else {
			// Handle other types of errors (e.g., JSON binding errors)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	ctx := c.Request.Context()
	id, err := p.Usecase.CreatePatient(ctx, &patientAdmin)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (p *PatientHandler) CreatePatientVisit(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id32 := int32(idP)
	ctx := c.Request.Context()

	var patientAdmin entities.Admin
	if err := c.ShouldBindJSON(&patientAdmin); err != nil {
		// Use type assertion to check if err is of type validator.ValidationErrors
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Get the first Validation Error
			fieldErr := validationErrs[0]
			c.JSON(http.StatusBadRequest, gin.H{"error": fieldErr.Error()})
			return // exit on first error
		} else if err.Error() == "EOF" {
			// Handle Empty Request Body Errors
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request Body is empty!"})
			return
		} else {
			// Handle other types of errors (e.g., JSON binding errors)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	vid, err := p.Usecase.CreatePatientVisit(ctx, id32, &patientAdmin)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vid": vid})
}

func (p *PatientHandler) DeletePatientVisit(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vidP, err := strconv.Atoi(c.Param("vid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id32 := int32(idP)
	vid32 := int32(vidP)
	ctx := c.Request.Context()

	err = p.Usecase.DeletePatientVisit(ctx, id32, vid32)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (p *PatientHandler) UpdatePatientVisit(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vidP, err := strconv.Atoi(c.Param("vid"))
	// Check if the id or vid is not a number
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id32 := int32(idP)
	vid32 := int32(vidP)
	ctx := c.Request.Context()

	var patient entities.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		// Use type assertion to check if err is of type validator.ValidationErrors
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Get the first Validation Error
			fieldErr := validationErrs[0]
			c.JSON(http.StatusBadRequest, gin.H{"error": fieldErr.Error()})
			return // exit on first error
		} else if err.Error() == "EOF" {
			// Handle Empty Request Body Errors
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request Body is empty!"})
			return
		} else {
			// Handle other types of errors (e.g., JSON binding errors)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	err = p.Usecase.UpdatePatientVisit(ctx, id32, vid32, &patient)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (p *PatientHandler) GetPatientMeta(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id32 := int32(idP)
	ctx := c.Request.Context()

	patientMeta, err := p.Usecase.GetPatientMeta(ctx, id32)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patientMeta)
}

func (p *PatientHandler) GetAllPatientVisitMeta(c *gin.Context) {
	dateStr := c.Param("date")

	ctx := c.Request.Context()

	var date time.Time
	var err error
	if dateStr == "default" {
		date = time.Time{}
	} else {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid date format: %s", dateStr)})
			return
		}
	}

	patientVisitMeta, err := p.Usecase.GetAllPatientVisitMeta(ctx, date)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patientVisitMeta)
}

func (p *PatientHandler) ExportDatabaseToCSV(c *gin.Context) {
	ctx := c.Request.Context()

	filePath := util.MustGitPath("repository/tmp/output.csv")
	err := p.Usecase.ExportDatabaseToCSV(ctx)
	if err != nil {
		log.Printf("Failed to export data to CSV: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to export data"})
		return
	}

	c.Writer.Header().Set("Content-Type", "text/csv")
	// Set the content disposition header to force download
	c.Writer.Header().Set("Content-Disposition", "attachment")

	// Write the contents of the CSV file to the response
	c.FileAttachment(filePath, "output.csv")
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
	case entities.ErrPatientVisitNotFound:
		return http.StatusNotFound
	case entities.ErrMissingAdminCategory:
		return http.StatusBadRequest
	case entities.ErrAuthenticationFailed:
		return http.StatusUnauthorized
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
