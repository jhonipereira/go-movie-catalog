package main

import (
	"flag"
	"fmt"
	"jhonidev/go/go-movie-catalog/internal/repository"
	"jhonidev/go/go-movie-catalog/internal/repository/dbrepo"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
	DSN    string //data source name
	DB     repository.DatabaseRepo
}

func main() {
	// set app config
	var app application

	// read from cmd
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to the DB
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close() //when the (main) func exits, it exits too

	//start webserver
	app.Domain = "example.com"
	log.Printf("Application started at port %d\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
