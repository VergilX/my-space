package db

type User struct {
	ID       int
	username string
	password string // hashed value
}

func (db DB) userSchemaExists() (bool, error) {

}

func (db DB) createUserTable() error {

}

func (db DB) dropUserTable() error {

}

func (db DB) addUser(user User) error {

}

func (db DB) updateUser(oldUsername string, newUser User) error {

}

func (db DB) deleteUser(username string) error {

}
