package repository

import (
	"Go_Courses/pkg/models"
	"testing"
)

const connStr = "postgres://postgres:123@localhost:5432/postgres"

func Test_BooksCRUD(t *testing.T) {
	db, err := New(connStr)
	if err != nil {
		t.Fatal(err.Error())
	}
	data, err := db.GetBooks()
	if err != nil {
		t.Fatal(err.Error())
	}
	lenght := len(data)
	book := models.Book{AuthorID: 16, GenreID: 1, Name: `Лох`}
	err = db.NewBook(book)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = db.DelBook(10)
	if err != nil {
		t.Fatal(err.Error())
	}
	data, err = db.GetBooks()
	if err != nil || lenght != len(data) {
		t.Fatal(err.Error())
	}
	_, err = db.GetID(1)
	if err != nil {
		t.Fatal(err.Error())
	}
	dataAuthor, err := db.GetAuthors()
	if err != nil {
		t.Fatal(err.Error())
	}
	lenght = len(dataAuthor)
	author := models.Author{Name: `Лох`}
	err = db.NewAuthor(author.Name)
	if err != nil {
		t.Fatal(err.Error())
	}
	dataAuthor, err = db.GetAuthors()
	if err != nil || lenght == len(dataAuthor) {
		t.Fatal(err.Error())
	}
	dataGenres, err := db.GetGenres()
	if err != nil {
		t.Fatal(err.Error())
	}
	lenght = len(dataGenres)
	_, err = db.NewGenre("author.Name")
	if err != nil {
		t.Fatal(err.Error())
	}
	dataGenres, err = db.GetGenres()
	if err != nil || lenght == len(dataGenres) {
		t.Fatal(err.Error())
	}
}
