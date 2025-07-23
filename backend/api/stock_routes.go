package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-test/db"
	"go-test/services"
)

func RegisterStockRoutes(router *gin.Engine) {
	router.GET("/stocks", func(c *gin.Context) {
		conn, err := db.ConnectDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection failed"})
			return
		}
		defer conn.Close()

		filters := map[string]string{}
		for _, key := range []string{"company", "brokerage", "ticker", "rating_from", "rating_to", "action", "sort_by", "order", "limit", "offset"} {
			if v := c.Query(key); v != "" {
				filters[key] = v
			}
		}

		stocks, err := db.ListStocks(conn, filters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, stocks)
	})

	router.GET("/recommendation", func(c *gin.Context) {
		conn, err := db.ConnectDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection failed"})
			return
		}
		defer conn.Close()

		topParam := c.DefaultQuery("top", "5")
		topN, err := strconv.Atoi(topParam)
		if err != nil || topN <= 0 {
			topN = 5
		}

		topStocks, err := services.RecommendTopStocks(conn, topN)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, topStocks)
	})

	router.GET("/rating-list/from", func(c *gin.Context) {
		conn, err := db.ConnectDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection failed"})
			return
		}
		defer conn.Close()

		ratingList, err := db.RatingFromList(conn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, ratingList)
	})

	router.GET("/rating-list/to", func(c *gin.Context) {
		conn, err := db.ConnectDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection failed"})
			return
		}
		defer conn.Close()

		ratingList, err := db.RatingToList(conn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, ratingList)
	})
}
