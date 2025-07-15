package db

import (
	"database/sql"
	"fmt"
	"go-test/models"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "postgresql://user:password@host:26257/dbname?sslmode=require"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Ping error:", err)
	}

	fmt.Println("Connected to CockroachDB")
}

func SaveData(items []models.StockInformation) {
	for _, item := range items {
		_, err := DB.Exec(`
			INSERT INTO stock_information (
				ticker,
				target_from,
				target_to,
				company,
				action,
				brokerage,
				rating_from,
				rating_to,
				time
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			item.Ticker,
			item.TargetFrom,
			item.TargetTo,
			item.Company,
			item.Action,
			item.Brokerage,
			item.RatingFrom,
			item.RatingTo,
			item.Time,
		)
		if err != nil {
			log.Println("Insert error:", err)
		}
	}
}
