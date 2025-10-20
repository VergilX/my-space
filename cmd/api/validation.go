package main

import (
	"github.com/VergilX/my-space/internal/dblayer"
	"github.com/VergilX/my-space/internal/validator"
)

func validateUser(v *validator.Validator, user *dblayer.CreateUserParams) bool {
	v.Check(user.Username == "", "username", "must not be empty")
	v.Check(user.Password == "", "password", "must not be empty")

	return v.Vaild()
}
