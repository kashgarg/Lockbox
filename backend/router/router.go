package router

import (
	"github.com/gorilla/mux"
	"github.com/kashgarg/lockbox/backend/handlers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/vaultitems", handlers.CreateVaultItem).Methods("POST")
	r.HandleFunc("/vaultitems", handlers.GetVaultItems).Methods("GET")

	return r
}
