package main

import (
	"Go_Courses/pkg/api"
	"Go_Courses/pkg/repository"
	"log"

	"github.com/gorilla/mux"
)

// const connStr = "postgres://postgres:123@localhost:5432/postgres"

func main() {
	db, err := repository.New("postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err.Error())
	}
	api := api.New(mux.NewRouter(), db)
	api.FillEndpoints()
	log.Fatal(api.ListenAndServe("localhost:8090"))
}
