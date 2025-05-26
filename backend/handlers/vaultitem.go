package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kashgarg/lockbox/backend/db"
	"github.com/kashgarg/lockbox/backend/models"
	"github.com/kashgarg/lockbox/backend/utils"
)

func CreateVaultItem(w http.ResponseWriter, r *http.Request) {
	var item models.VaultItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	encryptedPassword, err := utils.Encrypt(item.Password)
	if err != nil {
		http.Error(w, "Failed to encrypt password", http.StatusInternalServerError)
		fmt.Println("Encryption error:", err)
		return
	}
	item.Password = encryptedPassword

	query := `INSERT INTO vault_items (title, username, password, notes) 
	          VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err = db.Conn.QueryRow(context.Background(), query, item.Title, item.Username, item.Password, item.Notes).
		Scan(&item.ID, &item.CreatedAt)

	if err != nil {
		http.Error(w, "Failed to insert item", http.StatusInternalServerError)
		fmt.Println("Insert error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func GetVaultItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Conn.Query(context.Background(),
		`SELECT id, title, username, password, notes, created_at FROM vault_items`)
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		fmt.Println("Query error:", err)
		return
	}
	defer rows.Close()

	var items []models.VaultItem
	for rows.Next() {
		var item models.VaultItem
		err := rows.Scan(&item.ID, &item.Title, &item.Username, &item.Password, &item.Notes, &item.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to scan item", http.StatusInternalServerError)
			fmt.Println("Scan error:", err)
			return
		}

		decryptedPassword, err := utils.Decrypt(item.Password)
		if err != nil {
			http.Error(w, "Failed to decrypt password", http.StatusInternalServerError)
			fmt.Println("Decryption error:", err)
			return
		}
		item.Password = decryptedPassword


		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
}
