package main

import (
	"base-code-api/app/api"
	"base-code-api/app/db"
)

func main() {
	db.Connect()
	defer db.Close()
	api.StartServer()
}
