package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/fakercrud.db")
	if err != nil {
		return nil, fmt.Errorf("ERROR: Can not open sqlite database: %w", err)
	}

	return db, err
}

func CreateInitialTable() error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			full_name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			phone TEXT,
			address TEXT
		)`); err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	fmt.Println("Table users created successfully")

	return nil
}

func ReadUser(id int) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			full_name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			phone TEXT,
			address TEXT
		)`); err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	fmt.Println("Table users created successfully")

	return nil
}

func UpdateUser(id int, fullname string, email string, phone string, address string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	res, err := db.Exec(`
		UPDATE users
		SET full_name = ?, email = ?, phone = ?, address = ?
		WHERE id = ?
	`, id, fullname, email, phone, address)
	if err != nil {
		return fmt.Errorf("unable to update user %d: %w", id, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("unable to determine rows affected for user %d: %w", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}

	return nil
}

func DeleteUser(id int) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			full_name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			phone TEXT,
			address TEXT
		)`); err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	fmt.Println("Table users created successfully")

	return nil
}
