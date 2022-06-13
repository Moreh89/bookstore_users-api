package services

import (
	"github.com/Moreh89/bookstore_users-api/domain/users"
	"github.com/Moreh89/bookstore_users-api/utils/crypto_utils"
	"github.com/Moreh89/bookstore_users-api/utils/date_utils"
	"github.com/Moreh89/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(user users.User) (*users.User, *errors.RestError)
	GetUser(userId int64) (*users.User, *errors.RestError)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError)
	DeleteUser(userId int64) *errors.RestError
	Search(status string) (users.Users, *errors.RestError)
}

func (s *usersService)CreateUser(user users.User) (*users.User, *errors.RestError) {
	// if err := user.Validate(); err != nil {
	// 	return nil, err
	// }
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowString()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService)GetUser(userId int64) (*users.User, *errors.RestError) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService)UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if current.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if current.LastName != "" {
			current.LastName = user.LastName
		}
		if current.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService)DeleteUser(userId int64) *errors.RestError {
	// // current, err := GetUser(userId)
	// if err != nil {
	// 	return err
	// }
	// if err = current.Delete(); err != nil {
	// 	return err
	// }
	// return nil
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService)Search(status string) (users.Users, *errors.RestError) {
	// dao := &users.User{}
	// users, err := dao.FindByStatus(status)
	// if err != nil {
	// 	return nil, err
	// }
	// return users, nil
	dao := &users.User{}
	return dao.FindByStatus(status)
}
