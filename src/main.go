package main

import (
	"go-bitly/src/database"
	"go-bitly/src/database/migrations"
	"go-bitly/src/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.dev")

	database.Setup()

	migrations.RunMigrations()

	server.SetupAndListen()
}
