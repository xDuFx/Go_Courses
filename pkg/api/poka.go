package api

import (
	"Go_Courses/pkg/models"
	"encoding/json"
	"net/http"
)

func PokaHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
