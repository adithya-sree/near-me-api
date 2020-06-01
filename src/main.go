package main

import (
	"nearme-api/src/app"
	"nearme-api/src/config"
)

func main() {
	a := app.App{}
	config := config.GetAppConfig()
	err := a.Initialize(config)
	if err != nil {
		panic(err)
	}
	a.Run()
}
