package models

import "time"

type VaultItem struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}
