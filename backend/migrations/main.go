package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loadind .env file")
	}
}

func main() {
	RunMigrations()
}
