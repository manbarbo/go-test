package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("BD_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	options := os.Getenv("DB_OPTIONS")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s",
		user, pass, host, port, name, options)

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		log.Fatal("Missing one or more required DB environment variables: " + connStr)
		return nil, errors.New("missing DB env vars")
	}

	DB, err := sql.Open("pgx", connStr)
	fmt.Println(connStr)
	if err != nil {
		log.Fatal("Failed to connect:", err, connStr)
		return nil, err
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Ping error:", err)
		return nil, fmt.Errorf("DB not reachable: %w", err)
	}

	fmt.Println("Connected to CockroachDB")
	return DB, nil
}
