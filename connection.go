package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("ошибка при загрузке .env.local файла: %v\n", err)
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println("Параметры подключения: ", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}

	fmt.Println("Объект db после sql.Open: ", db)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Не удалось установить соединение с базой данных: %w", err)
	}

	return db, nil
}
