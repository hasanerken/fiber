package storages

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"os"
	"time"
)

func NewPostgreSQLConnection() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(os.Getenv("DATABASE_URL"))
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	// set some connection pool settings
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	boil.SetDB(db)
	boil.SetLocation(time.UTC)

	err = waitForDatabase(10, dsn)

	if err != nil {
		log.Fatalf("Error waiting for database: %v", err)
	}

	return db, nil
}

// Waiting for the database to be ready to accept connections.
func waitForDatabase(maxAttempts int, dsn string) error {
	for i := 0; i < maxAttempts; i++ {
		db, err := sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				err := db.Close()
				if err != nil {
					return err
				}
				return nil
			}
		}
		fmt.Printf("Attempt %d: waiting for database\n", i+1)
		time.Sleep(5 * time.Second)
	}
	return fmt.Errorf("unable to connect to database after %d attempts", maxAttempts)
}
