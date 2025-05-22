package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kashgarg/lockbox/backend/router"
)

func main() {
	router.SetupRoutes()

	port := ":8080"
	fmt.Printf("Server running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

