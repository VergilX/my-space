package db

import "database/sql"

var TABLENAME string = "Users"

type User struct {
	ID       int
	Username string
	Password string // hashed value
}

func (db *DB) UserTableExists() (bool, error) {
	return db.ifTableExists(TABLENAME)
}

func (db *DB) CreateUserTable() error {
	query := `
		CREATE TABLE $1(
			ID INT PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password TEXT
			)
		`
	_, err := db.Exec(query, TABLENAME)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DropUserTable() error {
	return db.dropTable(TABLENAME)
}

func (db *DB) AddUser(user User) error {
	query := `
		INSERT INTO $1 VALUES($2, $3, $4)
		`

	_, err := db.Exec(query, user.ID, user.Username, user.Password)

	return err
}

func (db *DB) IfUserExists(username string) (bool, error) {
	query := `SELECT * FROM $1 WHERE Username = $2`

	row := new(User)
	err := db.QueryRow(query, TABLENAME, username).Scan(&row)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

/*
func (db DB) UpdateUser(oldUsername string, newUser User) error {

}

func (db DB) DeleteUser(username string) error {

}

*/
