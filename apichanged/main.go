package main

import (
	"fmt"

	"apichanged/config"
	"apichanged/models"
	"apichanged/routing"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&models.Users{}, &models.Posts{})
	config.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Users{})
	config.DB.Model(&models.Posts{}).AddForeignKey("postid", "post(id)", "CASCADE", "CASCADE")

	router := routing.RouterMappings()

	//running
	router.Run()

}
