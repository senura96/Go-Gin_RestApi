package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB // DB global variable

//db configuration property declaration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		DBName:   "todos",
		Password: "neymar",
	}
	return &dbConfig
}

//setting the values
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc-Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}