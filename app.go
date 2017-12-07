// app.go

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App aaa.
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize create a database connection
func (a *App) Initialize(user, password, dbname, host, port string) {
	// connectionString :=
	// 	fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
	// 		user,
	// 		password,
	// 		dbname,
	// 		host,
	// 		port)
	// connectionString := fmt.Sprintf(
	// 	"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
	// 	"postgres",
	// 	"postgres",
	// 	os.Getenv("DB_ENV_POSTGRES_PASSWORD"),
	// os.Getenv("GOTESTRESTAPIWITHPOSTGRES_POSTGRES_1_PORT_5432_TCP_ADDR"),
	// os.Getenv("GOTESTRESTAPIWITHPOSTGRES_POSTGRES_1_PORT_5432_TCP_PORT"),
	// )

	// var err error
	// a.DB, err = sql.Open("postgres", connectionString)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	a.Router = mux.NewRouter()
}

// Run simply starts the application.
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
