package service

import (
	"errors"

	"github.com/deathcore666/authms/dbclient"
	"github.com/deathcore666/authms/model"
)

var errEmptyFields = errors.New("empty field entry")

//AuthService lol
type AuthService interface {
	Login(username, password string) (bool, error)
	Register(username, password string) (int, error)
}

type authService struct{}

func (authService) Login(username, password string) (bool, error) {
	if username == "" && password == "" {
		return false, errEmptyFields
	}
	var user model.UserAccount
	user.UserName = username
	user.Password = password
	if err := dbclient.QueryUser(user); err != nil {
		return false, err
	}
	return true, nil
}

func (authService) Register(username, password string) (id int, err error) {
	if username == "" && password == "" {
		return 0, errEmptyFields
	}
	var user model.UserAccount
	user.UserName = username
	user.Password = password
	id, err = dbclient.InsertUser(user)
	return
}
