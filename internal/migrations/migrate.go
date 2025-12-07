package migrations

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	migrations "github.com/lemeow125/goapi/internal/migrations/books"
	_ "modernc.org/sqlite"
)

func CreateDB() *sqlx.DB{
    // Create DB
    SQLITE_DB := os.Getenv("SQLITE_DB")
    db, err := sqlx.Connect("sqlite", SQLITE_DB)
    if err != nil {
        log.Fatal(err)
    }
    if SQLITE_DB == ""  {
        SQLITE_DB = "/tmp/db.sqlite3"
        log.Default().Println("No SQLite DB specified, using ephemeral directory" + SQLITE_DB)
    }

    return db
}

func Migrate(db *sqlx.DB) {
    // Run all migrations
    migrations.SetupBooks(db)
}

func Setup() *sqlx.DB {
    db := CreateDB()
    Migrate(db)
    return db
}
