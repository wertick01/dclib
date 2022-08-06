package processors

import (
	"errors"

	"github.com/wertick01/dclib/internals/app/db"
	"github.com/wertick01/dclib/internals/app/models"
)

type BooksProcessor struct {
	storage *db.BooksStorage
}

func NewBooksProcessor(storage *db.BooksStorage) *BooksProcessor {
	processor := new(BooksProcessor)
	processor.storage = storage
	return processor
}

func (processor *BooksProcessor) CreateBook(book *models.Books) (*models.Books, error) {

	if book.BookName == "" {
		return processor.storage.NullBooks(), errors.New("name should not be empty")
	}

	return processor.storage.CreateNewBook(book)
}

func (processor *BooksProcessor) FindBook(id int64) (*models.Books, error) {
	book, err := processor.storage.GetBookById(id)

	if err != nil {
		return processor.storage.NullBooks(), err
	}

	return book, nil

}

func (processor *BooksProcessor) ListBooks() ([]*models.Books, error) {
	list, err := processor.storage.GetBooksList()
	return list, err
}

func (processor *BooksProcessor) UpdateBook(book *models.Books) (*models.Books, error) { //!!! ПРОВЕРИТЬ
	return processor.storage.ChangeBook(book)
}

func (processor *BooksProcessor) DeleteBook(id int64) (int64, error) {
	deleted, err := processor.storage.DeleteBookById(id)
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (processor *BooksProcessor) StarTheBook(id int64) (*models.Books, error) {
	err := processor.storage.PutStarByBookId(id)
	if err != nil {
		return processor.storage.NullBooks(), err
	}
	book, err := processor.FindBook(id)
	if err != nil {
		return processor.storage.NullBooks(), err
	}
	return book, nil
}
