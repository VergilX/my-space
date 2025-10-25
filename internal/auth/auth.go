package auth

import (
	"errors"

	_ "github.com/VergilX/my-space/internal/db"
)

// temporary database
var users = map[string]string{}
var session = make(map[string]string)
var csrf = make(map[string]string)

// used for generating a token
var TOKEN_SIZE int = 128

func Register(username, password string) error {
	// check if username already in db
	if _, exists := users[username]; exists {
		return errors.New("username already exists")
	}

	// hash password
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}

	// store in db
	users[username] = hash

	// return error
	return nil
}

func Login(username, password string) (string, string, error) {
	// check if username in db
	hash, exists := users[username]
	if !exists {
		return "", "", errors.New("username does not exist")
	}

	// match username and password
	matched := ValidatePasswordWithHash(password, hash)
	if !matched {
		return "", "", errors.New("invalid password")
	}

	// create the session cookie
	session_token, err := GenerateToken(TOKEN_SIZE)
	if err != nil {
		return "", "", err
	}
	session[username] = session_token

	// create the csrf token, to be used by client for next request
	csrf_token, err := GenerateToken(TOKEN_SIZE)
	if err != nil {
		return "", "", err
	}

	// store in db
	csrf[username] = csrf_token

	return session_token, csrf_token, nil
}

func Logout(username string) error {
	// remove entry from db
	delete(users, username)
	delete(session, username)
	delete(csrf, username)

	return nil
}
