package repository

import (
	"Go_Courses/pkg/models"
	"context"
)

func (repo *PGRepo) GetBooks() ([]models.Book, error) {
	rows, err := repo.pool.Query(context.Background(), `
		    SELECT id, name, author_id, genre_id
FROM books;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []models.Book
	for rows.Next() {
		var item models.Book
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.AuthorID,
			&item.GenreID,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func (repo *PGRepo) NewBook(item models.Book) error {
	_, err := repo.pool.Exec(context.Background(), `
	INSERT INTO books(name, author_id, genre_id)
    VALUES ($1, $2, $3);`,
		item.Name,
		item.AuthorID,
		item.GenreID,
	)
	return err
}
func (repo *PGRepo) NewAuthor(name string) error {
	_, err := repo.pool.Exec(context.Background(), `
	INSERT INTO authors(name)
    VALUES ($1);`,
		name,
	)
	return err
}

func (repo *PGRepo) GetID(id int) (models.Book, error) {
	row := repo.pool.QueryRow(context.Background(), `
	SELECT id, name, author_id, genre_id
	FROM books 
	WHERE id = $1;`,
		id,
	)
	var data models.Book
	err := row.Scan(
		&data.ID,
		&data.Name,
		&data.AuthorID,
		&data.GenreID,
	)
	return data, err
}

func (repo *PGRepo) DelBook(id int) error {
	_, err := repo.pool.Exec(context.Background(), `
	DELETE from books
    WHERE id = $1;`,
		id,
	)
	return err
}

func (repo *PGRepo) GetAuthors() ([]models.Author, error) {
	rows, err := repo.pool.Query(context.Background(), `
		    SELECT id, name
FROM authors;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []models.Author
	for rows.Next() {
		var item models.Author
		err := rows.Scan(
			&item.ID,
			&item.Name,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}
