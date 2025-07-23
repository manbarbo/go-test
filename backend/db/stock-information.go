package db

import (
	"database/sql"
	"fmt"
	"go-test/models"
	"strings"
)

// InsertStock: Creates a stock information
func InsertStock(db *sql.DB, stock *models.StockInformation) error {
	err := db.QueryRow(`
		INSERT INTO stock_information (
			ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id`,
		stock.Ticker, stock.TargetFrom, stock.TargetTo, stock.Company,
		stock.Action, stock.Brokerage, stock.RatingFrom, stock.RatingTo, stock.Time).Scan(&stock.ID)
	return err
}

// ListStocks: Lists the stocks with filters
func ListStocks(db *sql.DB, filters map[string]string) ([]models.StockInformation, error) {
	query := `SELECT id, ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time FROM stock_information WHERE 1=1`
	args := []interface{}{}
	argPos := 1

	// Company Filter (includes)
	if v, ok := filters["company"]; ok && v != "" {
		query += fmt.Sprintf(" AND company ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Brokerage Filter (includes)
	if v, ok := filters["brokerage"]; ok && v != "" {
		query += fmt.Sprintf(" AND brokerage ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Ticker Filter (includes)
	if v, ok := filters["ticker"]; ok && v != "" {
		query += fmt.Sprintf(" AND ticker ILIKE $%d", argPos)
		args = append(args, "%"+v+"%")
		argPos++
	}

	// Action Filter (exact)
	if v, ok := filters["action"]; ok && v != "" {
		query += fmt.Sprintf(" AND action = $%d", argPos)
		args = append(args, v)
		argPos++
	}

	// Rating From Filter (exact)
	if v, ok := filters["rating_from"]; ok && v != "" {
		query += fmt.Sprintf(" AND rating_from = $%d", argPos)
		args = append(args, v)
		argPos++
	}

	// Rating To Filter (exact)
	if v, ok := filters["rating_to"]; ok && v != "" {
		query += fmt.Sprintf(" AND rating_to = $%d", argPos)
		args = append(args, v)
		argPos++
	}

	// Sort
	if sortBy, ok := filters["sort_by"]; ok && sortBy != "" {
		allowed := map[string]bool{
			"ticker": true, "company": true, "brokerage": true,
			"action": true, "time": true, "rating_from": true, "rating_to": true,
		}
		col := sortBy

		order := "ASC"
		if o, ok := filters["order"]; ok {
			upper := strings.ToUpper(o)
			if upper == "ASC" || upper == "DESC" {
				order = upper
			}
		}

		if allowed[col] && (order == "ASC" || order == "DESC") {
			query += fmt.Sprintf(" ORDER BY %s %s", col, order)
		}
	}

	// Pagination: limit y offset
	if limit, ok := filters["limit"]; ok && limit != "" {
		query += fmt.Sprintf(" LIMIT $%d", argPos)
		args = append(args, limit)
		argPos++
	}
	if offset, ok := filters["offset"]; ok && offset != "" {
		query += fmt.Sprintf(" OFFSET $%d", argPos)
		args = append(args, offset)
		argPos++
	}

	fmt.Println(query)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []models.StockInformation
	for rows.Next() {
		var s models.StockInformation
		err := rows.Scan(&s.ID, &s.Ticker, &s.TargetFrom, &s.TargetTo, &s.Company,
			&s.Action, &s.Brokerage, &s.RatingFrom, &s.RatingTo, &s.Time)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, s)
	}
	return stocks, nil
}

// GetStockByID: Gets an specific stock information by id
func GetStockByID(db *sql.DB, id string) (*models.StockInformation, error) {
	row := db.QueryRow(`
		SELECT id, ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time
		FROM stock_information WHERE id = $1`, id)

	var s models.StockInformation
	err := row.Scan(&s.ID, &s.Ticker, &s.TargetFrom, &s.TargetTo, &s.Company,
		&s.Action, &s.Brokerage, &s.RatingFrom, &s.RatingTo, &s.Time)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no encontrado
		}
		return nil, err
	}
	return &s, nil
}

// UpdateStock: Updates the stock information by id
func UpdateStock(db *sql.DB, stock models.StockInformation) error {
	_, err := db.Exec(`
		UPDATE stock_information SET
			ticker = $1,
			target_from = $2,
			target_to = $3,
			company = $4,
			action = $5,
			brokerage = $6,
			rating_from = $7,
			rating_to = $8,
			time = $9
		WHERE id = $10`,
		stock.Ticker, stock.TargetFrom, stock.TargetTo, stock.Company,
		stock.Action, stock.Brokerage, stock.RatingFrom, stock.RatingTo, stock.Time,
		stock.ID)
	return err
}

// DeleteStock: Deletes the stock by id
func DeleteStock(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM stock_information WHERE id = $1", id)
	return err
}

// ClearStocks: Deletes all the data
func ClearStocks(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM stock_information")
	return err
}

func RatingFromList(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT DISTINCT rating_from FROM stock_information ORDER BY rating_from")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []string
	for rows.Next() {
		var rating string
		if err := rows.Scan(&rating); err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}

func RatingToList(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT DISTINCT rating_to FROM stock_information ORDER BY rating_to")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []string
	for rows.Next() {
		var rating string
		if err := rows.Scan(&rating); err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}
