package database

import (
	"alfa-backend/modules"
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
	dbPort = getEnv("DB_PORT", "5433")
	dbUser = getEnv("DB_USER", "postgres")
	dbPassword = os.Getenv("DB_PASSWORD")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Connect(dbName string) (*sql.DB, error) {

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
		log.Println("CreateDatabase error:", err)
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
	}
	return nil
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

func AddUserToDB(db *sql.DB, email, password, name string) error {
	query := `INSERT INTO users (email, password, name) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, email, password, name)

	return err
}

func GetEmailFromDb(db *sql.DB, email string) (*modules.User, error) {
	var user modules.User
	query := `SELECT id, email, password, name, created_at FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Created_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("database error %v", err)
	}

	return &user, nil
}
