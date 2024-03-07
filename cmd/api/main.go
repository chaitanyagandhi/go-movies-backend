package main

import (
	"fmt"
	"go-movies-backend/internal/repository"
	"go-movies-backend/internal/repository/dbrepo"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	DSN          string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	APIKey       string
}

func main() {
	// err := godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	var app application

	app.DSN = os.Getenv("DB_CONNECTION_STRING")
	app.JWTSecret = os.Getenv("JWT_SECRET")
	app.JWTIssuer = "example.com"
	app.JWTAudience = "example.com"
	app.CookieDomain = "localhost"
	app.Domain = "example.com"
	app.APIKey = os.Getenv("API_KEY")

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.CookieDomain,
	}

	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), app.routes())

}
