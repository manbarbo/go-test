package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
}

func main() {
	// DB, err := db.ConnectDB()
	// if err != nil {
	// 	log.Fatalf("Error conectando a la base de datos: %v", err)
	// }

	// // Clear Stocks before insert
	// if err := db.ClearStocks(DB); err != nil {
	// 	log.Fatalf("Error limpiando tabla: %v", err)
	// }

	// // Load data and inster into the DB
	// err = api.LoadData(func(stock *models.StockInformation) error {
	// 	return db.InsertStock(DB, stock)
	// })
	// if err != nil {
	// 	log.Fatalf("Error cargando e insertando datos: %v", err)
	// }
}
