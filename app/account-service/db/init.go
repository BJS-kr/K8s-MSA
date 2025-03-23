package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Init(conn *sql.DB) error {
	_, err := conn.Exec(`
	CREATE DATABASE IF NOT EXISTS account_service
	`)

	if err != nil {
		return err
	}

	_, err = conn.Exec(`
	CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		name VARCHAR(10) NOT NULL
	)
	`)

	if err != nil {
		return err
	}

	_, err = conn.Exec(`
	CREATE TABLE IF NOT EXISTS balances (
		account_id INTEGER,
		currency INTEGER,
		balance INTEGER,

		PRIMARY KEY (account_id, currency),
		FOREIGN KEY (account_id) REFERENCES accounts(id)
	)
	`)

	if err != nil {
		return err
	}

	return nil
}
