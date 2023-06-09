package main

import (
	"flag"
	"fmt"
	_ "github.com/nanmenkaimak/restaurant-reservation-system/docs"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/handlers"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

const portNumber = ":8080"

// @title Restaurant Reservation System
// @version 1.0
// @description rest api in golang

// @host localhost:8080
// @BasePath /

// securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	content, err := os.ReadFile("password.txt")
	if err != nil {
		log.Fatal(err)
	}
	dbHost := flag.String("dbhost", "db", "Database host")
	dbName := flag.String("dbname", "restaurant-reservation-system", "Database name")
	dbUser := flag.String("dbuser", "postgres", "Database user")
	dbPass := flag.String("dbpass", string(content), "Database password")
	dbPort := flag.String("dbport", "5432", "Database port")
	dbSSL := flag.String("dbssl", "disable", "Database ssl settings (disable, prefer, require)")

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}))

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Users{}, &models.Roles{}, &models.Restaurants{}, &models.TypeOfRestaurant{}, &models.Food{}, &models.CategoryOfFood{}, &models.Reservations{})

	flag.Parse()

	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)

	routes()
}
