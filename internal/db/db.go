package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// creates an embedding which allows the dev
// to create new methods on it
type DB struct {
	*sql.DB
}

func New(dsn string) (*DB, error) {
	database, err := sql.Open("sqlite3", dsn)
	if err != nil {
		fmt.Println("database init error")
	}

	// use struct to create embedding
	embedding := DB{database}

	return &embedding, nil
}

func (db DB) Close() error {
	err := db.Close()
	if err != nil {
		fmt.Println("DB connection termination error")
		return err
	}

	return nil
}

func (db DB) createTable(tableName string) error {

}

func (db DB) dropTable(tableName string) error {

}

func (db DB) selectAllFromTable(tableName string) (*sql.Rows, error) {

}

func (db DB) selectRowFromTable(tableName string, key string, value string) (*sql.Row, error) {
	// idk if this function will work
}
