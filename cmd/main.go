package main

import (
	"Go_Courses/pkg/api"
	"Go_Courses/pkg/repository"
	"log"

	"github.com/gorilla/mux"
)


func main() {
	db, err := repository.New("postgres://Имя_пользователя:пароль@localhost:порт/название_БД")
	if err != nil {
		log.Fatal(err.Error())
	}
	api := api.New(mux.NewRouter(), db)
	api.FillEndpoints()
	log.Fatal(api.ListenAndServe("localhost:8090"))
}
