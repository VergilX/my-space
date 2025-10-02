package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	// cannot use embedding since using it in main
	// will not promote embedding methods to parent
	// struct
	*sql.DB
}

func New(dsn string) (*DB, error) {
	database, err := sql.Open("sqlite3", dsn)
	if err != nil {
		fmt.Println("database init error")
		return nil, err
	}

	// use struct
	db := DB{database}

	// create schema if it doesn't exist
	exists, err := db.UserTableExists()
	if err != nil {
		return nil, err
	}

	if !exists {
		err = db.CreateUserTable()
		if err != nil {
			return nil, err
		}
	}

	return &db, nil
}

func (db *DB) CloseConn() error {
	err := db.Close()
	if err != nil {
		fmt.Println("DB connection termination error")
		return err
	}

	return nil
}

func (db *DB) ifTableExists(tableName string) (bool, error) {
	query := `SELECT * FROM Users;`

	err := db.QueryRow(query, tableName).Scan()

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (db *DB) dropTable(tableName string) error {
	query := `DROP TABLE $1`

	_, err := db.Exec(query, tableName)
	if err != nil {
		return err
	}

	return nil
}

/*
func (db *DB) selectAllFromTable(tableName string) (*sql.Rows, error) {

}

func (db *DB) selectRowFromTable(tableName string, key string, value string) (*sql.Row, error) {
	// idk if this function will work
}

*/
