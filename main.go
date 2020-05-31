package main

import "nearme-api/app"

const connection = "root:APBlmKG4aIBg5xhI@tcp(35.238.6.1:3306)/locationdata"

func main() {
	a := app.App{}

	err := a.Initialize(connection)
	if err != nil {
		panic(err)
	}

	a.Run(":8888")
}
