package main

import (
	"antar-jemput/auth"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	e := echo.New()

	// Configuration du validateur
	e.Validator = &CustomValidator{validator.New()}

	// Chaîne de connexion sans mot de passe
	dsn := "host=localhost user=postgres dbname=antar-jemput port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	auth.RegisterAuthRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}

// CustomValidator est un wrapper pour le package de validation
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
