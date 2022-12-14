package db

import (
	"database/sql"

	"github.com/wertick01/dclib/internals/app/models"
)

type FavorietesStorage struct {
	AuthorsStorage
}

func NewFavorietesStorage(db *sql.DB) *FavorietesStorage {
	storage := new(FavorietesStorage)
	storage.DB = db
	return storage
}

func (m *FavorietesStorage) AddToFavorieteBooks(favorietes *models.FavorieteBooks) (int64, error) {

	stmt := `INSERT INTO dclib_test.favoriete_books (userid, book_id) VALUES(?, ?)`

	result, err := m.DB.Exec(stmt, favorietes.UserId, favorietes.FavoriteBookId)
	if err != nil {

		return 0, err
	}

	err = m.PutStarByBookId(favorietes.FavoriteBookId, "put")
	if err != nil {

		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {

		return 0, err
	}

	return id, nil
}

func (m *FavorietesStorage) AddToFavorieteAuthors(favorietes *models.FavorieteAuthors) (int64, error) {

	stmt := `INSERT INTO dclib_test.favoriete_authors (userid, author_id) VALUES(?, ?)`

	result, err := m.DB.Exec(stmt, favorietes.UserId, favorietes.FavoriteAuthorId)
	if err != nil {

		return 0, err
	}

	err = m.PutStarByAuthorId(favorietes.FavoriteAuthorId, "put")
	if err != nil {

		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {

		return 0, err
	}

	return id, nil
}

func (m *FavorietesStorage) GetFavoriteBooksList(user_id int64) ([]*models.Books, error) {

	stmt := `SELECT fb.book_id, b.book_name, b.book_count, b.book_photo, b.book_stars, ba.author_id, a.author_name, a.author_surname, a.author_patrynomic, a.author_photo, a.author_stars FROM dclib_test.favoriete_books AS fb RIGHT JOIN dclib_test.books AS b ON fb.book_id = b.book_id LEFT JOIN dclib_test.books_authors AS ba ON fb.book_id=ba.book_id INNER JOIN dclib_test.authors AS a ON ba.author_id=a.author_id WHERE fb.userid=?`

	rows, err := m.DB.Query(stmt, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmap []*models.Books

	for rows.Next() {
		s := &models.Books{}
		a := &models.Authors{}
		err = rows.Scan(&s.BookId, &s.BookName, &s.Count, &s.BookPhoto, &s.Stars, &a.AuthorId, &a.AuthorName.Name, &a.AuthorName.Surname, &a.AuthorName.Patronymic, &a.AuthorPhoto, &a.AuthorStars)
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

func (m *FavorietesStorage) GetFavoriteAuthorsList(user_id int64) ([]*models.Authors, error) {

	stmt := `SELECT fa.author_id, a.author_name, a.author_surname, a.author_patrynomic, a.author_photo, a.author_stars FROM dclib_test.favoriete_authors AS fa RIGHT JOIN dclib_test.authors AS a ON fa.author_id = a.author_id WHERE fa.userid = ?`

	var myauthors []*models.Authors

	rows, err := m.DB.Query(stmt, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		a := &models.Authors{}
		err = rows.Scan(&a.AuthorId, &a.AuthorName.Name, &a.AuthorName.Surname, &a.AuthorName.Patronymic, &a.AuthorPhoto, &a.AuthorStars)
		if err != nil {
			return nil, err
		}

		myauthors = append(myauthors, a)
	}
	return myauthors, nil
}

func (m *FavorietesStorage) DeleteFavorieteBookById(favoriete *models.FavorieteBooks) (int64, error) {
	stmt := `DELETE FROM dclib_test.favoriete_books WHERE book_id = ? AND userid = ?`

	deleted, err := m.DB.Exec(stmt, favoriete.FavoriteBookId, favoriete.UserId)
	if err != nil {

		return 0, err
	}

	err = m.PutStarByBookId(favoriete.FavoriteBookId, "delete")
	if err != nil {

		return 0, err
	}

	res, err := deleted.LastInsertId()
	if err != nil {

		return 0, err
	}
	return res, nil
}

func (m *FavorietesStorage) DeleteFavorieteAuthorById(favoriete *models.FavorieteAuthors) (int64, error) {
	stmt := `DELETE FROM dclib_test.favoriete_authors WHERE author_id = ? AND userid = ?`

	deleted, err := m.DB.Exec(stmt, favoriete.FavoriteAuthorId, favoriete.UserId)
	if err != nil {

		return 0, err
	}

	err = m.PutStarByAuthorId(favoriete.FavoriteAuthorId, "delete")
	if err != nil {

		return 0, err
	}

	res, err := deleted.LastInsertId()
	if err != nil {

		return 0, err
	}
	return res, nil
}
