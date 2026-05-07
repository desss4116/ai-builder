package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open(
		"sqlite3",
		"knowledge.db",
	)

	if err != nil {
		log.Fatal(err)
	}

	optimizeSQLite()

	createTables()
}

func optimizeSQLite() {

	pragmas := []string{

		"PRAGMA journal_mode=WAL;",

		"PRAGMA synchronous=OFF;",

		"PRAGMA temp_store=MEMORY;",

		"PRAGMA cache_size=-20000;",

		"PRAGMA mmap_size=30000000000;",
	}

	for _, p := range pragmas {

		DB.Exec(p)
	}
}

func createTables() {

	query := `
	CREATE VIRTUAL TABLE IF NOT EXISTS knowledge_base
	USING fts5(
		category,
		title,
		content
	);
	`

	_, err := DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}
