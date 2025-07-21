package services

import (
	"database/sql"
	"sort"
	"strconv"
	"strings"
	"time"

	"go-test/db"
	"go-test/models"
	"go-test/utils"
)

func RecommendTopStocks(sdb *sql.DB, topN int) ([]models.StockWithScore, error) {
	stocks, err := db.ListStocks(sdb, map[string]string{})
	if err != nil {
		return nil, err
	}

	var scored []models.StockWithScore
	for _, s := range stocks {
		score := calculateScore(s)
		scored = append(scored, models.StockWithScore{StockInformation: s, Score: score})
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	if topN > len(scored) {
		topN = len(scored)
	}

	return scored[:topN], nil
}

func calculateScore(stock models.StockInformation) float64 {
	score := 0.0

	from := utils.RatingScore[stock.RatingFrom]
	to := utils.RatingScore[stock.RatingTo]
	score += float64(to-from) * 2.5

	fromPrice := parsePrice(stock.TargetFrom)
	toPrice := parsePrice(stock.TargetTo)
	if fromPrice > 0 && toPrice > fromPrice {
		score += ((toPrice - fromPrice) / fromPrice) * 100
	}

	if t, err := time.Parse(time.RFC3339, stock.Time); err == nil {
		daysAgo := time.Since(t).Hours() / 24
		if daysAgo < 30 {
			score += (30 - daysAgo) * 0.1
		}
	}

	return score
}

func parsePrice(price string) float64 {
	cleaned := strings.Replace(price, "$", "", -1)
	cleaned = strings.TrimSpace(cleaned)
	val, _ := strconv.ParseFloat(cleaned, 64)
	return val
}
