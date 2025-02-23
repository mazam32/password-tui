package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const dbFile = "passwords.db"

type DBManager struct {
	DB *sql.DB
}

func NewDBManager() (*DBManager, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS passwords (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		service TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &DBManager{DB: db}, nil
}

func (db *DBManager) AddPassword(service string, password string) error {
	_, err := db.DB.Exec("INSERT INTO passwords (service, password) VALUES (?, ?)", service, password)

	if err != nil {
		return fmt.Errorf("failed to add password: %w", err)
	}

	return nil
}

func (db *DBManager) GetPassword(service string) (string, error) {
	var password string
	err := db.DB.QueryRow("SELECT password FROM passwords WHERE service = ?", service).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no password found for service: %s", service)
		}
		return "", err
	}
	return password, nil
}

func (db *DBManager) ListServices() ([]string, error) {
	rows, err := db.DB.Query("SELECT service FROM passwords")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var services []string
	for rows.Next() {
		var service string
		if err := rows.Scan(&service); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (db *DBManager) RemovePassword(service string) error {
	_, err := db.DB.Exec("DELETE FROM passwords WHERE service = ?", service)
	if err != nil {
		return fmt.Errorf("failed to remove password: %w", err)
	}
	return nil
}
