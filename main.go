package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

const (
	SHUFFLE_INTERVAL = 24 * time.Hour // Shuffle interval set to 48 hours
)

// shuffleTable shuffles the specified table by updating the order_col with random values
func shuffleTable(db *sql.DB, tableName string) {
	query := fmt.Sprintf("UPDATE %s SET order_col = random()", tableName)
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error shuffling table %s: %v", tableName, err)
	} else {
		log.Printf("Table %s shuffled successfully", tableName)
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection parameters
	PostgresDatabase := os.Getenv("POSTGRES_DATABASE")
	PostgresUser := os.Getenv("POSTGRES_USER")
	PostgresPassword := os.Getenv("POSTGRES_PASSWORD")
	PostgresHost := os.Getenv("POSTGRES_HOST")

	// Open the database connection
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s host=%s password=%s dbname=%s sslmode=disable",
		PostgresUser, PostgresHost, PostgresPassword, PostgresDatabase,
	))
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Tables to be shuffled
	tables := []string{"banklogs", "cashapp", "credit_cards"}

	// Infinite loop to shuffle tables periodically
	for {
		// Log the current time of shuffle
		currentTime := time.Now()
		log.Printf("Shuffling started at: %s", currentTime.Format("2006-01-02 15:04:05"))

		// Shuffle each table
		for _, table := range tables {
			shuffleTable(db, table)
		}

		// Calculate and log the next shuffle time
		nextShuffle := currentTime.Add(SHUFFLE_INTERVAL)
		log.Printf("Shuffling tables completed. Next shuffle scheduled at: %s\n",
			nextShuffle.Format("2006-01-02 15:04:05"))

		// Sleep for the specified interval before the next shuffle
		time.Sleep(SHUFFLE_INTERVAL)
	}
}
