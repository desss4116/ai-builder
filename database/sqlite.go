package database

import (
	"database/sql"
	_ "://github.com"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./ai_builder.db")
	if err != nil {
		log.Fatal(err)
	}

	// Таблица для "знаний" системы (Knowledge Base)
	query := `
	CREATE TABLE IF NOT EXISTS knowledge (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tag TEXT,
		content TEXT,
		source TEXT
	);
	CREATE TABLE IF NOT EXISTS sites (
		id TEXT PRIMARY KEY,
		prompt TEXT,
		html TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

