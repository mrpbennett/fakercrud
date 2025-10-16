// Package db provides helpers for managing the fakercrud SQLite database.
package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mrpbennett/fakercrud/generateuser"
)

// DBConnection establishes and returns a connection to the SQLite database.
// It opens a connection to the fakercrud.db file located in the ./db directory.
// Returns a pointer to sql.DB and an error if the connection fails.
// The caller is responsible for closing the database connection when done.
func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/fakercrud.db")
	if err != nil {
		return nil, fmt.Errorf("ERROR: Can not open sqlite database: %w", err)
	}

	return db, err
}

// CreateInitialTable creates the initial users table in the SQLite database if it doesn't exist.
// The table includes columns for id (auto-increment primary key), full_name, email (unique),
// phone, and address. It establishes a database connection, creates the table, and closes
// the connection automatically.
// Returns an error if the database connection fails or table creation fails.
func CreateInitialTable() error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`
		create table if not exists users (
			id integer primary key autoincrement,
			full_name text not null,
			email text not null unique,
			phone text,
			address text
		)`); err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	return nil
}

// InitialiseUsers populates the users table with n randomly generated fake users.
// It generates fake user data using the generateuser package and inserts them in a
// single transaction for efficiency. If any insertion fails, the entire transaction
// is rolled back to maintain data integrity. Upon successful completion, it prints
// the number of users inserted to stdout.
// The n parameter specifies the number of users to generate and insert.
// Returns an error if database connection, transaction, or insertion fails.
func InitialiseUsers(n int) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	users := generateuser.GenerateUsers(n)

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("unable to start insert transaction: %w", err)
	}

	stmt, err := tx.Prepare(`
		insert into users (full_name, email, phone, address)
		values (?, ?, ?, ?)
	`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("unable to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	for _, user := range users {
		if _, err := stmt.Exec(user.FullName, user.Email, user.Phone, user.Address); err != nil {
			tx.Rollback()
			return fmt.Errorf("unable to insert user %q: %w", user.Email, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("unable to commit inserted users: %w", err)
	}

	fmt.Printf("%d users inserted successfully\n", len(users))

	return nil
}

func CreateUser(fullname string, email string, phone string, address string) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// res, err := db.Exec(`
	// 	update users
	// 	set full_name = ?, email = ?, phone = ?, address = ?
	// 	where id = ?
	// `,  fullname, email, phone, address)
	// if err != nil {
	// 	return fmt.Errorf("unable to update user %d: %w",  err)
	// }

	return nil
}

// ReturnUser looks up a single user by id and returns the hydrated domain model.
// It returns sql.ErrNoRows wrapped in context when the user does not exist.
func ReturnUser(id int) (*generateuser.User, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user generateuser.User

	row := db.QueryRowContext(context.Background(), `
		select full_name, email, phone, address
		from users
		where id = ?
	`, id)

	if err := row.Scan(&user.FullName, &user.Email, &user.Phone, &user.Address); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user %d not found: %w", id, err)
		}
		return nil, fmt.Errorf("querying user %d: %w", id, err)
	}

	return &user, nil
}

// UpdateUser persists the provided field values for the user identified by id
// and returns the updated user. Consumers should handle the wrapped sql.ErrNoRows
// error when the user cannot be found.
func UpdateUser(id int, fullname string, email string, phone string, address string) (*generateuser.User, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user generateuser.User

	// ExecContext issues the write-only UPDATE, mirroring database/sql's intended usage.
	result, err := db.ExecContext(context.Background(), `
		update users
		set full_name = ?, email = ?, phone = ?, address = ?
		where id = ?`, fullname, email, phone, address, id)
	if err != nil {
		return nil, fmt.Errorf("updating user %d: %w", id, err)
	}

	// Capture RowsAffected to signal callers when the target id does not exist.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("checking update result for user %d: %w", id, err)
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("user %d not found: %w", id, sql.ErrNoRows)
	}

	// Follow-up SELECT reloads the freshly updated row for the caller.
	row := db.QueryRowContext(context.Background(), `
		select full_name, email, phone, address
		from users
		where id = ?`, id)

	if err := row.Scan(&user.FullName, &user.Email, &user.Phone, &user.Address); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user %d not found after update: %w", id, err)
		}
		return nil, fmt.Errorf("retrieving updated user %d: %w", id, err)
	}

	return &user, nil
}

// DeleteUser removes the user identified by id. If the user does not exist the
// call still succeeds, mirroring standard SQL semantics.
func DeleteUser(id int) error {
	db, err := DBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`
		create table if not exists users (
			id integer primary key autoincrement,
			full_name text not null,
			email text not null unique,
			phone text,
			address text
		)`); err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	return nil
}
