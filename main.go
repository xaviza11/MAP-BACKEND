package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-sqlite-backend/database"
	"go-sqlite-backend/handlers"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Error loading PORT")
	}
	database.InitDB()
	defer database.DB.Close()

	http.HandleFunc("/user/create", handlers.CreateUserHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/country/create", handlers.CreateCountryHandler)

	fmt.Printf("Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
