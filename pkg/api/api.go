package api

import (
	"Go_Courses/pkg/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	r  *mux.Router
	db *repository.PGRepo
}

func New(router *mux.Router, db *repository.PGRepo) *api {
	return &api{r: router, db: db}
}

func (api *api) FillEndpoints() {
	api.r.HandleFunc("/api/books", api.books)
	api.r.HandleFunc("/api/author", api.authorNew).Queries("name", "{name}")
	api.r.HandleFunc("/api/author", api.authorNew)
	api.r.HandleFunc("/api/dBook", api.delete_book).Queries("id", "{id}")
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
