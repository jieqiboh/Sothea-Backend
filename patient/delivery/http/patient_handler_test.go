package http

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_patientPostgresRepository "github.com/jieqiboh/sothea_backend/patient/repository/postgres"
	_patientUseCase "github.com/jieqiboh/sothea_backend/patient/usecase"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func initServer() *gin.Engine {
	// todo: Convert to a mock, and add debug configs
	// Initialize global variables
	viper.SetConfigFile(`../../../config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbName := viper.GetString(`database.name`)
	dbSslMode := viper.GetString(`database.sslmode`)

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbName, dbSslMode)

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// You might want to check the connection here to handle errors
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()
	patientRepo := _patientPostgresRepository.NewPostgresPatientRepository(db)
	patientUseCase := _patientUseCase.NewPatientUsecase(patientRepo, 2*time.Second)
	NewPatientHandler(router, patientUseCase)
	return router
}

func TestPingRoute(t *testing.T) {
	router := initServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetPatientByID(t *testing.T) {
	router := initServer()

	jsonBody := `
	{
		"admin": {
			"familyGroup": "S001",
			"regDate": "2024-01-10T00:00:00Z",
			"name": "John Doe",
			"dob": "1994-01-10T00:00:00Z",
			"age": 30,
			"gender": "M",
			"village": "SO",
			"contactNo": "12345678",
			"pregnant": false,
			"drugAllergies": "panadol",
			"sentToID": true
		}
	}
	`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/patient", bytes.NewReader([]byte(jsonBody)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
