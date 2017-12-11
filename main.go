package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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

func serveIndex(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(resp, "Hello, World! 5")

	fmt.Fprintln(resp, "DB_ADDR:", os.Getenv("DB_PORT_5432_TCP_ADDR"))
	fmt.Fprintln(resp, "DB_PORT:", os.Getenv("DB_PORT_5432_TCP_PORT"))

	_, err := a.DB.Exec("insert into mydata(val) values(0)")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := a.DB.Query("select id from mydata")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(resp, "ID: %d\n", id)
	}
}
