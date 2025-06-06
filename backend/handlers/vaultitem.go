package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

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

func UpdateVaultItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

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

	// Perform the update
	query := `UPDATE vault_items SET title=$1, username=$2, password=$3, notes=$4 WHERE id=$5`
	_, err = db.Conn.Exec(context.Background(), query, item.Title, item.Username, encryptedPassword, item.Notes, id)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
		fmt.Println("Update error:", err)
		return
	}

	// Fetch the updated item
	query = `SELECT id, title, username, password, notes, created_at FROM vault_items WHERE id=$1`
	err = db.Conn.QueryRow(context.Background(), query, id).Scan(
		&item.ID, &item.Title, &item.Username, &item.Password, &item.Notes, &item.CreatedAt,
	)
	if err != nil {
		http.Error(w, "Failed to fetch updated item", http.StatusInternalServerError)
		fmt.Println("Fetch after update error:", err)
		return
	}

	// Decrypt password before sending to client
	item.Password, err = utils.Decrypt(item.Password)
	if err != nil {
		http.Error(w, "Failed to decrypt password", http.StatusInternalServerError)
		fmt.Println("Decryption error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func DeleteVaultItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	query := `DELETE FROM vault_items WHERE id=$1`

	_, err := db.Conn.Exec(context.Background(), query, id)
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
		fmt.Println("Delete error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Item deleted successfully")
}
