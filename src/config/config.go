package config

import (
	"fmt"
)

//AppConfig configs
type AppConfig struct {
	DbHost   string
	DbName   string
	Username string
	Password string
	AppPort  string
}

//GetAppConfig gets the application configs
func GetAppConfig() AppConfig {
	return AppConfig{
		DbHost:   "35.238.6.1:3306",
		DbName:   "locationdata",
		Username: "root",
		Password: "APBlmKG4aIBg5xhI",
		AppPort:  ":8888",
	}
}

//ConnectionString gets the DB connection stirng
func (a AppConfig) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", a.Username, a.Password, a.DbHost, a.DbName)
}
