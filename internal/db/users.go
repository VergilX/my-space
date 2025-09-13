package db

type User struct {
	ID       int
	username string
	password string // hashed value
}

func (db DB) userSchemaExists(tableName string) (bool, error) {

}

func (db DB) addUser(user User) error {

}

func (db DB) updateUser(oldUsername string, newUser User) error {

}

func (db DB) deleteUser(username string) error {

}
