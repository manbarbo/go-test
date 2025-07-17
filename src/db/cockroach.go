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
	connStr := os.Getenv("DB_CONN_STR")

	if connStr == "" {
		err := "DB_CONN_STR is not defined in .env file"
		log.Fatal(err)
		return nil, errors.New(err)
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
