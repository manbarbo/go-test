package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	connStr := os.Getenv("DB_CONN_STR_MIGRATIONS")
	if connStr == "" {
		log.Fatal("DB_CONN_STR is not defined in .env file")
	}

	m, err := migrate.New("file://migrations", connStr)
	if err != nil {
		log.Fatal("Error creating migrator:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error executiong migrations:", err)
	} else if err == migrate.ErrNoChange {
		log.Println("There are no new migrations to be applied.")
	} else {
		log.Println("Migrations executed successfully.")
	}
}
