package users

import (
	// "github.com/Moreh89/bookstore_users-api/utils/errors"
	// "strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"`
}

type Users []User

// function
// func Validate(user *User) (*errors.RestError){
//	// validation of email
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("invalid email address")
// 	}
// 	return nil
// }

// method inside user
// func (user *User) Validate() *errors.RestError {
// 	// validation of email
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("invalid email address")
// 	}
// 	return nil
// }
