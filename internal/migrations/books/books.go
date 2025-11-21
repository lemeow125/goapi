package migrations

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)


func SetupBooks(db *sqlx.DB) {
    query := `
    CREATE TABLE IF NOT EXISTS Books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title text NOT NULL CHECK(length(title) > 0 AND length(title) < 64), 
        author text NOT NULL CHECK(length(author) > 0 AND length(author) < 64)
    );
    `

    // Execute Query
    _, err := db.Exec(query)
    if err != nil {
        log.Fatalf("%q: %s\n", err, query)
        return
    }
}