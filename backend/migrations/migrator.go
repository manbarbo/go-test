package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("BD_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	options := os.Getenv("DB_OPTIONS")
	connector := os.Getenv("BD_MIGRATION_CONNECTOR")

	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?%s",
		connector, user, pass, host, port, name, options)

	if user == "" || pass == "" || host == "" || port == "" || name == "" || connector == "" {
		log.Fatal("Missing one or more required DB environment variables:" + connStr)
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
