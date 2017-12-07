package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	_ = App{}

	connInfo := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_ENV_POSTGRES_USER"),
		os.Getenv("DB_ENV_POSTGRES_DBNAME"),
		os.Getenv("DB_ENV_POSTGRES_PASSWORD"),
		os.Getenv("DB_PORT_5432_TCP_ADDR"),
		os.Getenv("DB_ENV_POSTGRES_PORT"),
	)

	var err error
	db, err = sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(i) * time.Second)

		if err = db.Ping(); err == nil {
			break
		}
		log.Println(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`create table if not exists mydata (
			id serial primary key,
			val integer not null
		)`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", serveIndex)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func serveIndex(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(resp, "Hello, World! 5")

	fmt.Fprintln(resp, "DB_ADDR:", os.Getenv("DB_PORT_5432_TCP_ADDR"))
	fmt.Fprintln(resp, "DB_PORT:", os.Getenv("DB_PORT_5432_TCP_PORT"))

	_, err := db.Exec("insert into mydata(val) values(0)")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id from mydata")
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
