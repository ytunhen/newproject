package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() error {
	var err error
	connStr := "user=ytunhen password=10040108Chuu dbname=test sslmode=disable host=db port=5432"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		processed BOOLEAN DEFAULT FALSE
	);`
	_, err = db.Exec(createTableQuery)
	return err
}

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

func saveMessageToDB(msg Message) error {
	_, err := db.Exec("INSERT INTO messages (content) VALUES ($1)", msg.Content)
	return err
}

func getProcessedMessagesStats() (map[string]int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM messages WHERE processed = TRUE").Scan(&count)
	if err != nil {
		return nil, err
	}
	return map[string]int{"processed_messages": count}, nil
}
