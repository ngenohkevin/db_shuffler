package main

import (
	"database/sql"
	"fmt"
	"log"
)

func shuffleTable(db *sql.DB, tableName string) {
	query := fmt.Sprintf("UPDATE %s SET order_col = random()", tableName)
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error shuffling table %s: %v", tableName, err)
	} else {
		log.Printf("Table %s shuffled Successfully", tableName)
	}
}

func main() {

}
