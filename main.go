package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_httpDelivery "github.com/jieqiboh/sothea_backend/controllers"
	_patientPostgresRepository "github.com/jieqiboh/sothea_backend/repository/postgres"
	_useCase "github.com/jieqiboh/sothea_backend/usecases"
	"github.com/spf13/viper"
	"log"
	"time"
)

const (
	isProduction = false // Global variable, change to true when in production mode!
)

func main() {
	// Initialize global variables
	if isProduction {
		viper.SetConfigFile(`prod.json`)
	} else {
		viper.SetConfigFile(`config.json`)
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	address := viper.GetString(`server.address`)
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbName := viper.GetString(`database.name`)
	dbPassword := viper.GetString(`database.password`)
	dbSslMode := viper.GetString(`database.sslmode`)
	secretKey := []byte(viper.GetString(`jwt.secretkey`))

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, dbSslMode)

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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PATCH"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Set up login routes
	loginUseCase := _useCase.NewLoginUseCase(5*time.Second, secretKey)
	_httpDelivery.NewLoginHandler(router, loginUseCase, secretKey)

	// Set up patient routes
	patientRepo := _patientPostgresRepository.NewPostgresPatientRepository(db)
	patientUseCase := _useCase.NewPatientUsecase(patientRepo, 2*time.Second)
	_httpDelivery.NewPatientHandler(router, patientUseCase, secretKey)

	router.Run(address)
}
