package migrations

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)


func SetupBooks(db *sqlx.DB) {
    query := `
    CREATE TABLE IF NOT EXISTS Books (
        id INTEGER AUTO_INCREMENT NOT NULL,
        title text NOT NULL, 
        author text NOT NULL,
        PRIMARY KEY (id)
    );
    `

    // Execute Query
    _, err := db.Exec(query)
    if err != nil {
        log.Fatalf("%q: %s\n", err, query)
        return
    }
}