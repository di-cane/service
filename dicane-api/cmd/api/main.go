package main

import (
	"database/sql"
	"dicane-api/data"
	"fmt"
	"log"
	"net/http"
	"os"
)

const webPort = "8000"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	// connect to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// set up config
	app := Config{
		DB:     db,
		Models: data.New(db),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	error := srv.ListenAndServe()
	if error != nil {
		log.Panic(err)
	}
}
