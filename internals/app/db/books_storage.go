package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/wertick01/dclib/internals/app/models"
)

type BooksStorage struct {
	DB *sql.DB
}

func NewBooksStorage(db *sql.DB) *BooksStorage {
	storage := new(BooksStorage)
	storage.DB = db
	return storage
}

func (m *BooksStorage) CreateNewBook(book *models.Books) (*models.Books, error) {

	stmt := `INSERT INTO dclib_test.books (book_name, book_count, book_photo) VALUES(?, ?, ?)`

	result, err := m.DB.Exec(stmt, book.BookName, book.Count, book.BookPhoto)
	if err != nil {
		return m.NullBooks(), err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return m.NullBooks(), err
	}
	fmt.Printf("---> Book %v has been added to DB", id)

	connected, err := m.BooksAuthorsConnection(id, book)
	if err != nil {
		return m.NullBooks(), err
	}

	fmt.Printf("---> Book %v has been added to DB", id)

	return connected, nil
}

func (m *BooksStorage) BooksAuthorsConnection(id int64, books *models.Books) (*models.Books, error) {
	stmt := `INSERT INTO dclib_test.books_authors (book_id, author_id) VALUES(?, ?)`
	//sdmd := `SELECT book_id FROM dclib_test.books WHERE book_name = ?`
	//res := m.DB.QueryRow(sdmd, )

	authors := books.Authors
	for _, val := range authors {
		row, err := m.DB.Exec(stmt, id, val.AuthorId)
		if err != nil {
			return m.NullBooks(), err
		}
		id, err := row.LastInsertId()
		if err != nil {
			return m.NullBooks(), err
		}
		fmt.Printf("Id %v has beed added to DB.", id)
	}
	return books, nil
}

/*
func (m *BooksStorage) GetBooksList() ([]*models.Books, error) {
	stmt := `SELECT  ba.book_id,  b.book_name,  b.book_count, b.book_photo, ba.author_id,  a.author_name, a.author_surname, a.author_patrynomic, a.author_photo
	FROM dclib_test.books_authors AS ba INNER JOIN dclib_test.books AS b USING(book_id) INNER JOIN dclib_test.authors AS a USING(author_id)`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allbooks map[int]*models.Books

	var a_id int64 = 0

	for rows.Next() {
		s := &models.Books{}
		a := models.Authors{}
		err = rows.Scan(&s.BookId,
			&s.BookName,
			&s.Count,
			&s.BookPhoto,

			&a.AuthorId,
			&a.AuthorName.Name,
			&a.AuthorName.Surname,
			&a.AuthorName.Patronymic,
			&a.AuthorPhoto,
		)
		if err != nil {
			return nil, err
		}

		if a.AuthorId != a_id {
			s.Authors = append(s.Authors, a)
		}
		a_id = a.AuthorId
	}

}
*/

func (m *BooksStorage) GetBooksList() ([]*models.Books, error) {

	stmt := `SELECT book_id, book_name, book_count FROM dclib_test.books`
	skmk := `SELECT author_id FROM dclib_test.books_authors WHERE book_id = ?`
	sdmd := `SELECT author_name, author_surname, author_patrynomic FROM dclib_test.authors WHERE author_id = ?`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allbooks []*models.Books

	for rows.Next() {
		var a_id []int64
		var id int64
		s := &models.Books{}
		err = rows.Scan(&s.BookId, &s.BookName, &s.Count)
		if err != nil {
			return nil, err
		}

		connection, err := m.DB.Query(skmk, s.BookId)
		if err != nil {
			return nil, err
		}

		defer connection.Close()

		for connection.Next() {
			err = connection.Scan(&id)
			if err != nil {
				return nil, err
			}

			a_id = append(a_id, id)
		}

		for _, val := range a_id {
			a := &models.Authors{}
			a.AuthorId = val
			authors := m.DB.QueryRow(sdmd, val)

			err = authors.Scan(
				&a.AuthorName.Name,
				&a.AuthorName.Surname,
				&a.AuthorName.Patronymic,
			)

			if err != nil {
				return nil, err
			}
			s.Authors = append(s.Authors, *a)
		}

		allbooks = append(allbooks, s)
	}

	return allbooks, nil
}

func (m *BooksStorage) GetBookById(id int64) (*models.Books, error) {

	stmt := `SELECT book_id, book_name, book_count FROM dclib_test.books WHERE book_id = ?`
	skmk := `SELECT author_id FROM dclib_test.books_authors WHERE book_id = ?`
	sdmd := `SELECT author_name, author_surname, author_patrynomic FROM dclib_test.authors WHERE author_id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Books{}

	var a_id []int64
	var id_ int64

	err := row.Scan(&s.BookId, &s.BookName, &s.Count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return m.NullBooks(), models.ErrNoRecord
		} else {
			return m.NullBooks(), err
		}
	}

	connection, err := m.DB.Query(skmk, s.BookId)
	if err != nil {
		return m.NullBooks(), err
	}

	defer connection.Close()

	for connection.Next() {
		err = connection.Scan(&id_)
		if err != nil {
			return m.NullBooks(), err
		}

		a_id = append(a_id, id_)
	}

	for _, val := range a_id {
		a := &models.Authors{}
		a.AuthorId = val
		authors := m.DB.QueryRow(sdmd, val)

		err = authors.Scan(
			&a.AuthorName.Name,
			&a.AuthorName.Surname,
			&a.AuthorName.Patronymic,
		)

		if err != nil {
			return m.NullBooks(), err
		}
		s.Authors = append(s.Authors, *a)
	}

	return s, nil
}

func (m *BooksStorage) ChangeBook(old *models.Books) (*models.Books, error) { //доделать с учётом нескольких авторов

	stmt := `UPDATE dclib_test.books SET book_name = ?, book_count = ?, book_photo = ? WHERE book_id = ?`
	sdmd := `DELETE FROM dclib_test.books_authors WHERE book_id = ?`

	change, err := m.DB.Exec(stmt, old.BookName, old.Count, old.BookPhoto, old.BookId)
	if err != nil {
		return m.NullBooks(), err
	}
	id_1, err := change.LastInsertId()
	if err != nil {
		return m.NullBooks(), err
	}
	fmt.Printf("Book %v has been changed.", id_1)

	deleted, err := m.DB.Exec(sdmd, old.BookId)
	if err != nil {
		return m.NullBooks(), err
	}
	id_2, err := deleted.LastInsertId()
	if err != nil {
		return m.NullBooks(), err
	}
	fmt.Println(id_2)

	connected, err := m.BooksAuthorsConnection(old.BookId, old)
	if err != nil {
		return m.NullBooks(), err
	}

	return connected, nil
}

func (m *BooksStorage) NullBooks() *models.Books {
	return &models.Books{
		BookId:    0,
		BookName:  "NO_NAME",
		BookPhoto: "NO_PHOTO",
	}
}

func (m *BooksStorage) DeleteBookById(id int64) (int64, error) {
	stmt := `DELETE FROM dclib_test.books WHERE book_id = ?`
	sdmd := `DELETE FROM dclib_test.books_authors WHERE book_id = ?`

	deleted, err := m.DB.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	res, err := deleted.LastInsertId()
	if err != nil {
		return 0, err
	}

	delet, err := m.DB.Exec(sdmd, id)
	if err != nil {
		return 0, err
	}

	result, err := delet.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Printf("Book %v has been deleted", result)
	return res, nil
}

func (m *BooksStorage) PutStarByBookId(id int64) error { //!!!

	stmt := `UPDATE dclib_test.books SET book_stars = ? WHERE book_id = ?`

	book, err := m.GetBookById(id)
	if err != nil {
		return err
	}

	book.Stars += 1

	putstar, err := m.DB.Exec(stmt, book.Stars, id)
	if err != nil {
		return err
	}

	id, err = putstar.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
