package models

import (
	"apichanged/config"
	//Driver Import
	_ "github.com/go-sql-driver/mysql"
)

//Getting All Users Details

func GetAllUserDetails(user *[]Users) (err error) {
	if err := config.DB.Find(user).Error; err != nil {

		return err
	}
	return nil

}

//Inserting a User
func CreateUser(user *Users) (err error) {
	if err := config.DB.Create(user).Error; err != nil {

		return err
	}

	return nil
}

// Retrieving a User Detail By Id

func GetAUser(user *Users, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(user).Error; err != nil {

		return err
	}

	return nil

}

// Updating user details

func UpdateAUser(user *Users, id string) (err error) {
	config.DB.Update(user)
	return nil

}

//Deleteing user details
func DeleteAUser(user *Users, id string) (err error) {

	config.DB.Where("id = ?", id).Delete(user)
	return nil

}
