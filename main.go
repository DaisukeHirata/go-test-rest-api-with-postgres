package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	a  App
)

func main() {
	a = App{}
	a.Initialize(
		os.Getenv("DB_ENV_POSTGRES_USER"),
		os.Getenv("DB_ENV_POSTGRES_DBNAME"),
		os.Getenv("DB_ENV_POSTGRES_PASSWORD"),
		os.Getenv("DB_PORT_5432_TCP_ADDR"),
		os.Getenv("DB_ENV_POSTGRES_PORT"),
	)

	a.Run(":8080")
}
