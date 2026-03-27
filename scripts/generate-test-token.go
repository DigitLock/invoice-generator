package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Try loading .env from CWD (run from backend/) or from project root
	_ = godotenv.Load(".env")
	_ = godotenv.Load("backend/.env")

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev-secret-change-in-production"
	}

	claims := jwt.MapClaims{
		"user_id":   1,
		"family_id": 1,
		"email":     "test@example.com",
		"name":      "Test User",
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatalf("Failed to sign token: %v", err)
	}

	fmt.Println(signed)
}
