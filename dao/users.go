package dao

import (
	"fmt"
	"github.com/jhonnli/go-base/model"
)

func GetUserList() []model.Users {
	var userList []model.Users
	defer db.Close()
	if err := db.Find(&userList).Error; err != nil {
		fmt.Println(err)
		return nil
	} else {
		return userList
	}
}
