package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jieqiboh/sothea_backend/domain"
	"net/http"
	"strconv"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// PatientHandler represent the httphandler for patient
type PatientHandler struct {
	AUsecase domain.PatientUseCase
}

// NewPatientHandler will initialize the patients/ resources endpoint
func NewPatientHandler(e *gin.Engine, us domain.PatientUseCase) {
	handler := &PatientHandler{
		AUsecase: us,
	}
	e.GET("/patient/:id", handler.GetPatientByID)
	e.POST("/patient", handler.InsertPatient)
	e.DELETE("/patient/:id", handler.DeletePatientByID)
	e.PATCH("/patient/:id", handler.UpdatePatientByID)
	e.GET("/patient/admin", handler.GetAllAdmin)
}

func (p *PatientHandler) GetPatientByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
		return
	}

	id := int64(idP)
	ctx := c.Request.Context()

	patient, err := p.AUsecase.GetPatientByID(ctx, id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (p *PatientHandler) InsertPatient(c *gin.Context) {
	var patient domain.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	ctx := c.Request.Context()
	id, err := p.AUsecase.InsertPatient(ctx, &patient)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "Inserted userid": id})
}

func (p *PatientHandler) DeletePatientByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
		return
	}

	id := int64(idP)
	ctx := c.Request.Context()

	id, err = p.AUsecase.DeletePatientByID(ctx, id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func (p *PatientHandler) UpdatePatientByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request.Context()

	var patient domain.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	id, err = p.AUsecase.UpdatePatientByID(ctx, id, &patient)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func (p *PatientHandler) GetAllAdmin(c *gin.Context) {
	ctx := c.Request.Context()
	arrAdmin, err := p.AUsecase.GetAllFromAdmin(ctx)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, arrAdmin)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
