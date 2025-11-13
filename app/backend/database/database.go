package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
)

func init() {
	loadConfig()
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	dbHost = getEnv("DB_HOST", "localhost")
	dbPort = getEnv("DB_PORT", "5432")
	dbUser = getEnv("DB_USER", "postgres")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = getEnv("DB_NAME", "postgres")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Connect() (*sql.DB, error) {

	connstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

func CreateDatabase(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE myapp")
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
	}
	return nil
}

func ConnectToDb() (*sql.DB, error) {
	connstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=myapp sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword)

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal("Error while connecting to myapp")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err

}

func CreateTables(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL, 
        name VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)

	return err
}
