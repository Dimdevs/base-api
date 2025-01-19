package main

import (
	"base-code-api/internal/api"
	"base-code-api/internal/db"
)

func main() {
	db.Connect()
	defer db.Close()
	api.StartServer()
}
