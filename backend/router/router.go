package router

import (
	"net/http"

	"github.com/kashgarg/lockbox/backend/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/health", handlers.HealthCheck)
}
