package main

import (
	"context"
	"database/sql"
)

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitMigrations(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS RATES (BID DECIMAL(10, 2))`)
	if err != nil {
		return err
	}

	return nil
}

func SaveExchangeRate(db *sql.DB, bid string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, DatabaseRateInsertTimeout)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO RATES (bid) VALUES (?)", bid)
	if err != nil {
		return nil
	}

	return nil
}
