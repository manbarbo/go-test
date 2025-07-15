package main

import (
	"fmt"
	"log"

	"go-test/api"
	"go-test/utils"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
}

func main() {
	// Llamar a la función LoadData
	data, err := api.LoadData()
	if err != nil {
		log.Fatalf("Error cargando los datos: %v", err)
	}

	utils.PrintDataAsJSON(data)
	fmt.Printf("Total datos: %d\n", len(data))
}
