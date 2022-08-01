package db

import (
	"database/sql"
	"fmt"

	"github.com/wertick01/dclib/internals/app/models"
)

type FavorietesStorage struct {
	DB *sql.DB
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
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Printf("---> book %v has been added to Favorietes", id)

	return id, nil
}

func (m *FavorietesStorage) AddToFavorieteAuthors(favorietes *models.FavorieteAuthors) (int64, error) {

	stmt := `INSERT INTO dclib_test.favoriete_authors (userid, author_id) VALUES(?, ?)`

	result, err := m.DB.Exec(stmt, favorietes.UserId, favorietes.FavoriteAuthorId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Printf("---> Author %v has been added to Favorietes\n", id)

	return id, nil
}

func (m *FavorietesStorage) GetFavoriteBooksList(user_id int64) ([]*models.Books, error) {

	stmt := `SELECT fb.book_id, b.book_name, b.book_count FROM dclib_test.favoriete_books AS fb RIGHT JOIN dclib_test.books AS b ON fb.book_id = b.book_id WHERE fb.userid = ?`
	skmk := `SELECT author_id FROM dclib_test.books_authors WHERE book_id = ?`
	sdmd := `SELECT author_name, author_surname, author_patrynomic FROM dclib_test.authors WHERE author_id = ?`

	var mybooks []*models.Books

	rows, err := m.DB.Query(stmt, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := &models.Books{}
		err = rows.Scan(&s.BookId, &s.BookName, &s.Count)
		if err != nil {
			return nil, err
		}

		bookstr, err := m.DB.Query(skmk, s.BookId)
		if err != nil {
			return nil, err
		}
		defer bookstr.Close()

		for bookstr.Next() {
			b := &models.Authors{}
			err = bookstr.Scan(&b.AuthorId)
			if err != nil {
				return nil, err
			}

			connection := m.DB.QueryRow(sdmd, b.AuthorId)

			err = connection.Scan(&b.AuthorName.Name, &b.AuthorName.Surname, &b.AuthorName.Patronymic)
			if err != nil {
				return nil, err
			}
			s.Authors = append(s.Authors, *b)

			mybooks = append(mybooks, s)
		}
	}
	return mybooks, nil
}

func (m *FavorietesStorage) GetFavoriteAuthorsList(user_id int64) ([]*models.Authors, error) {

	stmt := `SELECT fa.author_id, a.author_name, a.author_surname, a.author_patrynomic FROM dclib_test.favoriete_authors AS fa RIGHT JOIN dclib_test.authors AS a ON fa.author_id = a.author_id WHERE fa.userid = ?`

	var myauthors []*models.Authors

	rows, err := m.DB.Query(stmt, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		a := &models.Authors{}
		err = rows.Scan(&a.AuthorId, &a.AuthorName.Name, &a.AuthorName.Surname, &a.AuthorName.Patronymic)
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

	res, err := deleted.LastInsertId()
	if err != nil {
		return 0, err
	}
	return res, nil
}
