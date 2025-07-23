package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-test/api"
	"go-test/middlewares"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
}

func main() {
	router := gin.Default()

	corsMiddleware, _ := middlewares.CORSMiddleware()
	router.Use(corsMiddleware)

	api.RegisterLifeCheckRoutes(router)
	api.RegisterStockRoutes(router)

	router.Run(":8080")
}
