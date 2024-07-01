package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connectDb() (*sql.DB, error) {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("ошибка при загрузке .env.local файла: %v\n", err)
	}
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("строка подключения к базе данных не задана")
	}

	// Подключитесь к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}

	return db, nil
}
