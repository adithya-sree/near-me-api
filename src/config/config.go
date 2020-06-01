package config

import (
	"fmt"
)

//AppConfig configs
type AppConfig struct {
	DbHost      string
	DbName      string
	DBUsername  string
	DBPassword  string
	AppUsername string
	AppPassword string
	AppPort     string
}

//GetAppConfig gets the application configs
func GetAppConfig() AppConfig {
	return AppConfig{
		DbHost:      "35.238.6.1:3306",
		DbName:      "locationdata",
		DBUsername:  "root",
		DBPassword:  "APBlmKG4aIBg5xhI",
		AppUsername: "test",
		AppPassword: "test",
		AppPort:     ":8888",
	}
}

//ConnectionString gets the DB connection stirng
func (a AppConfig) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", a.DBUsername, a.DBPassword, a.DbHost, a.DbName)
}
