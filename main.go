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

func main() {
	// Initialize global variables
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbName := viper.GetString(`database.name`)
	dbPassword := viper.GetString(`database.password`)
	dbSslMode := viper.GetString(`database.sslmode`)

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
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Set up login routes
	loginUseCase := _useCase.NewLoginUseCase(5 * time.Second)
	_httpDelivery.NewLoginHandler(router, loginUseCase)

	// Set up patient routes
	patientRepo := _patientPostgresRepository.NewPostgresPatientRepository(db)
	patientUseCase := _useCase.NewPatientUsecase(patientRepo, 2*time.Second)
	_httpDelivery.NewPatientHandler(router, patientUseCase)

	router.Run("localhost:9090")
}
