package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect() error {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return fmt.Errorf("DATABASE_URL not set")
	}

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
	return nil
}
