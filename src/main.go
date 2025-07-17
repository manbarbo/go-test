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
		log.Fatal("Error cargando el archivo .env")
	}
}

func main() {
	DB, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	stockTest := models.StockInformation{
		Ticker:     "CECO",
		TargetFrom: "$33.00",
		TargetTo:   "$33.00",
		Company:    "CECO Environmental",
		Action:     "upgraded by",
		Brokerage:  "HC Wainwright",
		RatingFrom: "Neutral",
		RatingTo:   "Buy",
		Time:       "2025-05-01T00:30:06.015697838Z",
	}

	err = db.InsertStock(DB, &stockTest)
	if err != nil {
		log.Fatalf("Error insertando el stock: %v", err)
	} else {
		log.Println("Stock insertado exitosamente.")
	}
	// Llamar a la función LoadData
	/*data, err := api.LoadData()
	if err != nil {
		log.Fatalf("Error cargando los datos: %v", err)
	}

	utils.PrintDataAsJSON(data)
	fmt.Printf("Total datos: %d\n", len(data))*/
}
