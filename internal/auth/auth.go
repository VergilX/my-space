package auth

// import db

func Register() error {
	// hash password

	// store in db

	// return error
}

func Login(username, password string) (string, string, error) {
	// hash password

	// get username, password from db

	// match username and password

	// create the session cookie

	// create the csrf token, to be used by client for next request

}

func Logout(username string) error {
	// get username from db

	// remove entry
}
