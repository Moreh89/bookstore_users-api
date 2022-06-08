package users

import (
	"fmt"
	"strings"

	"github.com/Moreh89/bookstore_users-api/datasources/mysql/users_db"
	"github.com/Moreh89/bookstore_users-api/utils/date_utils"
	"github.com/Moreh89/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
)

func (user *User) Get() *errors.RestError {
	err := users_db.Client.Ping()
	if err != nil {
		panic(err)
	}
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError("error when trying to save user: %s", err.Error())
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error when trying to save user: %s", err.Error())
	}
	user.Id = userId
	return nil
}
