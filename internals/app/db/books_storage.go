package db

import (
	"database/sql"
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

	connected, err := m.BooksAuthorsConnection(id, book)
	if err != nil {
		return m.NullBooks(), err
	}

	return connected, nil
}

func (m *BooksStorage) BooksAuthorsConnection(id int64, books *models.Books) (*models.Books, error) {
	stmt := `INSERT INTO dclib_test.books_authors (book_id, author_id) VALUES(?, ?)`

	authors := books.Authors
	for _, val := range authors {
		row, err := m.DB.Exec(stmt, id, val.AuthorId)
		if err != nil {
			return m.NullBooks(), err
		}
		_, err = row.LastInsertId()
		if err != nil {
			return m.NullBooks(), err
		}
	}
	return books, nil
}

func (m *BooksStorage) GetBooksList() ([]*models.Books, error) {

	stmt := `SELECT ba.book_id, b.book_name, b.book_count, b.book_photo, ba.author_id, a.author_name, a.author_surname, a.author_patrynomic, a.author_photo FROM books_authors AS ba LEFT JOIN authors AS a ON ba.author_id=a.author_id RIGHT JOIN books AS b ON ba.book_id=b.book_id`

	rows, err := m.DB.Query(stmt)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bookmap []*models.Books

	for rows.Next() {
		s := &models.Books{}
		a := &models.Authors{}
		err = rows.Scan(&s.BookId, &s.BookName, &s.Count, &s.BookPhoto, &a.AuthorId, &a.AuthorName.Name, &a.AuthorName.Surname, &a.AuthorName.Patronymic, &a.AuthorPhoto)
		fmt.Println(err)
		if err != nil {
			return nil, models.ErrNoRecord
		}

		s.Authors = append(s.Authors, *a)
		bookmap = append(bookmap, s)

	}

	var a, b_id int64 = 0, 0
	var book *models.Books
	var books []*models.Books

	for _, val := range bookmap {

		if val.BookId == b_id {
			book.Authors = append(book.Authors, val.Authors[0])
			a--
			books = remove(books, a)
		} else {
			book = val
			b_id = val.BookId
		}
		a++
		books = append(books, book)
	}

	return books, nil
}

func (m *BooksStorage) GetBookById(id int64) (*models.Books, error) {

	stmt := `SELECT ba.book_id, b.book_name, b.book_count, b.book_photo, ba.author_id, a.author_name, a.author_surname, a.author_patrynomic, a.author_photo FROM books_authors AS ba LEFT JOIN authors AS a ON ba.author_id=a.author_id RIGHT JOIN books AS b ON ba.book_id=b.book_id WHERE ba.book_id=?`

	rows, err := m.DB.Query(stmt, id)
	if err != nil {
		return nil, models.ErrNoRecord
	}

	defer rows.Close()

	var bookmap []*models.Books
	var bookres = new(models.Books)
	var b int = 0

	for rows.Next() {
		s := &models.Books{}
		a := &models.Authors{}
		err = rows.Scan(&s.BookId, &s.BookName, &s.Count, &s.BookPhoto, &a.AuthorId, &a.AuthorName.Name, &a.AuthorName.Surname, &a.AuthorName.Patronymic, &a.AuthorPhoto)
		if err != nil {
			return nil, err
		}
		s.Authors = append(s.Authors, *a)
		bookmap = append(bookmap, s)
		bookres.Authors = append(bookres.Authors, *a)
		if b == 0 {
			bookres.BookId = s.BookId
			bookres.BookName = s.BookName
			bookres.Count = s.Count
			bookres.BookPhoto = s.BookPhoto
			bookres.Stars = s.Stars
		}
		b++
	}

	return bookres, nil
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

	delet, err := m.DB.Exec(sdmd, id)
	if err != nil {
		return 0, err
	}

	result, err := delet.LastInsertId()
	if err != nil {
		return 0, err
	}

	deleted, err := m.DB.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	res, err := deleted.LastInsertId()
	if err != nil {
		return 0, err
	}

	auth, err := m.GetBooksByAuthorId(id)
	if err != nil {
		return 0, err
	}

	if len(auth) < 1 {
		_, err := m.DeleteAuthorById(id)
		if err != nil {
			return 0, err
		}
	}

	fmt.Printf("Book %v has been deleted", result)
	return res, nil
}

func (m *BooksStorage) GetBooksByAuthorId(id int64) ([]*models.Books, error) {
	stmt := `SELECT book_id FROM dclib_test.books_authors WHERE author_id = ?`

	rows, err := m.DB.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	var book_id int64
	var books []*models.Books

	for rows.Next() {
		err = rows.Scan(&book_id)
		book, err := m.GetBookById(book_id)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (m *BooksStorage) DeleteAuthorById(id int64) (int64, error) {
	stmt := `DELETE FROM dclib_test.authors WHERE author_id = ?`
	sdmd := `DELETE FROM dclib_test.books_authors WHERE author_id = ?`

	_, err := m.DB.Exec(sdmd, id)
	if err != nil {
		return 0, err
	}

	deleted, err := m.DB.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	res, err := deleted.LastInsertId()
	if err != nil {
		return 0, err
	}

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

func remove(slice []*models.Books, s int64) []*models.Books {
	return append(slice[:s], slice[s+1:]...)
}
