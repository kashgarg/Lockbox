package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kashgarg/lockbox/backend/db"
	"github.com/kashgarg/lockbox/backend/router"
)

func main() {
	godotenv.Load()
	err := db.Connect()
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	r := router.SetupRoutes()

	port := ":8080"
	fmt.Printf("Server running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
