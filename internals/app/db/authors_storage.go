package db

import (
	"database/sql"
	"errors"

	"github.com/wertick01/dclib/internals/app/models"
)

type AuthorsStorage struct {
	BooksStorage
}

func NewAuthorsStorage(db *sql.DB) *AuthorsStorage {
	storage := new(AuthorsStorage)
	storage.DB = db
	return storage

}

func (m *AuthorsStorage) CreateNewAuthor(author *models.Authors) (*models.Authors, error) {

	stmt := `INSERT INTO dclib_test.authors (author_name, author_surname, author_patrynomic, author_photo) VALUES(?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, author.AuthorName.Name, author.AuthorName.Surname, author.AuthorName.Patronymic, author.AuthorPhoto)
	if err != nil {

		return nil, err
	}

	_, err = result.LastInsertId()
	if err != nil {

		return nil, err
	}

	return author, nil
}

func (m *AuthorsStorage) GetAuthorsList() ([]*models.Authors, error) {

	stmt := `SELECT author_id, author_name, author_surname, author_patrynomic, author_photo, author_stars FROM dclib_test.author`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err

	}

	defer rows.Close()

	var authors []*models.Authors

	for rows.Next() {
		s := &models.Authors{}
		err = rows.Scan(&s.AuthorId, &s.AuthorName.Name, &s.AuthorName.Surname, &s.AuthorName.Patronymic, &s.AuthorPhoto, &s.AuthorStars)

		if err != nil {

			return nil, err
		}
		authors = append(authors, s)
	}

	if err = rows.Err(); err != nil {

		return nil, err
	}

	return authors, nil
}

func (m *AuthorsStorage) NullAuthors() *models.Authors {
	return &models.Authors{
		AuthorId: 0,
		AuthorName: models.AuthorName{
			Name:       "NO NAME",
			Surname:    "NO SURNAME",
			Patronymic: "NO PATRYNOMIC",
		},
	}
}

func (m *AuthorsStorage) GetAuthorById(id int64) (*models.Authors, error) {

	stmt := `SELECT author_id, author_name, author_surname, author_patrynomic, author_photo, author_stars FROM dclib_test.authors WHERE author_id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Authors{}

	err := row.Scan(&s.AuthorId, &s.AuthorName.Name, &s.AuthorName.Surname, &s.AuthorName.Patronymic, &s.AuthorPhoto, &s.AuthorStars)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {

			return nil, err

		} else {

			return nil, err
		}
	}

	return s, nil
}

func (m *AuthorsStorage) GetBooksByAuthorId(id int64) ([]*models.Books, *models.Authors, error) {
	stmt := `SELECT book_id FROM dclib_test.books_authors WHERE author_id = ?`

	rows, err := m.DB.Query(stmt, id)

	if err != nil {

		return nil, m.NullAuthors(), err
	}
	var book_id int64
	var books []*models.Books

	for rows.Next() {

		err = rows.Scan(&book_id)
		if err != nil {

			return nil, m.NullAuthors(), err
		}

		book, err := m.GetBookById(book_id)
		if err != nil {

			return nil, m.NullAuthors(), err
		}

		books = append(books, book)
	}
	author, err := m.GetAuthorById(id)
	if err != nil {

		return nil, m.NullAuthors(), err
	}

	return books, author, nil
}

func (m *AuthorsStorage) PutStarByAuthorId(id int64, method string) error { //!!!

	stmt := `UPDATE dclib_test.authors SET author_stars = ? WHERE author_id = ?`

	author, err := m.GetAuthorById(id)
	if errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if method == "put" {
		author.AuthorStars += 1
	}
	if method == "delete" {
		author.AuthorStars -= 1
	}

	putstar, err := m.DB.Exec(stmt, author.AuthorStars, id)
	if err != nil {
		return err
	}

	id, err = putstar.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (m *AuthorsStorage) ChangeAuthor(old *models.Authors) (*models.Authors, error) {

	stmt := `UPDATE dclib_test.authors SET author_name = ?, author_surname = ?, author_patrynomic = ?, author_photo = ? WHERE author_id = ?`
	sdmd := `DELETE FROM dclib_test.books_authors WHERE author_id = ?`

	change, err := m.DB.Exec(stmt, old.AuthorName.Name, old.AuthorName.Surname, old.AuthorName.Patronymic, old.AuthorPhoto, old.AuthorId)
	if err != nil {
		return m.NullAuthors(), err
	}
	_, err = change.LastInsertId()
	if err != nil {
		return m.NullAuthors(), err
	}

	deleted, err := m.DB.Exec(sdmd, old.AuthorId)
	if err != nil {
		return m.NullAuthors(), err
	}
	_, err = deleted.LastInsertId()
	if err != nil {
		return m.NullAuthors(), err
	}

	connected, err := m.AuthorsBooksConnection(old)
	if err != nil {
		return m.NullAuthors(), err
	}

	return connected, nil
}

func (m *AuthorsStorage) AuthorsBooksConnection(author *models.Authors) (*models.Authors, error) {
	stmt := `INSERT INTO dclib_test.books_authors (book_id, author_id) VALUES(?, ?)`

	books, author, err := m.GetBooksByAuthorId(author.AuthorId)
	if err != nil {
		return m.NullAuthors(), err
	}
	for _, val := range books {
		row, err := m.DB.Exec(stmt, val.BookId, author.AuthorId)
		if err != nil {
			return m.NullAuthors(), err
		}
		_, err = row.LastInsertId()
		if err != nil {
			return m.NullAuthors(), err
		}
	}
	return author, nil
}
