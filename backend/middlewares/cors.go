package middlewares

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() (gin.HandlerFunc, error) {
	origins := os.Getenv("CORS_ORIGINS")
	if origins == "" {
		log.Fatal("Missing CORS_ORIGINS environment variable")
		return nil, errors.New("missing DB env vars")
	}
	log.Println(origins)
	allowedOrigins := strings.Split(origins, ",")

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}), nil
}
