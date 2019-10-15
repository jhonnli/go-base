package users

import (
	"github.com/jhonnli/go-base/dao"
	"github.com/jhonnli/go-base/model"
)

func GetUsers() (users []model.Users) {
	users = dao.GetUserList()
	return users
}
