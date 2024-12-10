package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Open a database connection
	db, err := sql.Open("mysql", "admin:BaLiDiNeSh5@tcp(database-1.c1w2ua0029vr.us-east-2.rds.amazonaws.com:3306)/doordash?multiStatements=true")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Create a database driver instance
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Error creating database driver: %v", err)
	}

	// Create a new migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Use "file://" prefix for local migrations
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Error creating migrate instance: %v", err)
	}

	// Run migrations (change steps as needed)
	err = m.Steps(7)
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No changes to apply")
		} else {
			log.Fatalf("Error applying migrations: %v", err)
		}
	}

	// Close the migrate instance
	sourceErr, dbErr := m.Close()
	if sourceErr != nil || dbErr != nil {
		log.Printf("Error closing migration resources: sourceErr=%v, dbErr=%v", sourceErr, dbErr)
	}

	log.Println("Migrations applied successfully")
}
