package models

import (
	"apichanged/config"
	//Driver Import
	_ "github.com/go-sql-driver/mysql"
)

func FindUserByUNPW(user *Users, username string, password string) (err error) {
	if err := config.DB.Where("username = ? AND password = ?", username, password).Find(user).Error; err != nil {

		return err

	}

	return nil

}
