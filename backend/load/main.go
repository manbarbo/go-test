package main

import (
	"go-test/db"
	"go-test/models"
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

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Limpiar la tabla antes de insertar
	if err := db.ClearStocks(dbConn); err != nil {
		log.Fatalf("Error clearing table: %v", err)
	}

	// Pasar la función que inserta
	err = LoadData(func(stock *models.StockInformation) error {
		return db.InsertStock(dbConn, stock)
	})
	if err != nil {
		log.Fatalf("Error loading and inserting data: %v", err)
	}

	log.Println("Data upload completed successfully.")
}
