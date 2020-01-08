package models

import (
	"apichanged/config"
	//Driver Import
	_ "github.com/go-sql-driver/mysql"
)

//Getting All Posts Details

func GetAllPosts(post *[]Posts) (err error) {
	if err := config.DB.Find(post).Error; err != nil {

		return err
	}

	return nil
}

//Getting All Posts By User Id
func GetPostByUserId(posts *[]Posts, userId string) (err error) {

	if err := config.DB.Where("userid=?", userId).Find(posts).Error; err != nil {

		return err
	}

	return nil

}

//Inserting A Post

func CreateAPost(post *Posts) (err error) {

	if err := config.DB.Create(post).Error; err != nil {

		return err
	}

	return nil

}

//Retrieving A Post Detail By Post Id

func GetAPost(post *Posts, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(post).Error; err != nil {

		return err
	}
	return nil

}

//Updating a post

func UpdateAPost(post *Posts, id string) (err error) {
	config.DB.Update(post)

	return nil
}

//Deleting a post

func DeleteAPost(post *Posts, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(post)

	return nil

}
