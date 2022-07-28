package processors

import (
	"errors"
	"fmt"

	"github.com/wertick01/dclib/internals/app/db"
	"github.com/wertick01/dclib/internals/app/models"
)

type AuthorsProcessor struct {
	storage *db.AuthorsStorage
}

func NewAuthorsProcessor(storage *db.AuthorsStorage) *AuthorsProcessor {
	processor := new(AuthorsProcessor)
	processor.storage = storage
	return processor
}

func (processor *AuthorsProcessor) CreateAuthor(author *models.Authors) (int64, error) {

	if author.AuthorName.Name == "" && author.AuthorName.Surname == "" {
		return 0, errors.New("name should not be empty")
	}

	return processor.storage.CreateNewAuthor(author)
}

func (processor *AuthorsProcessor) ListAuthors() ([]*models.Authors, error) {
	return processor.storage.GetAuthorsList()
}

func (processor *AuthorsProcessor) AuthorsBooks(id int64) ([]*models.Books, *models.Authors, error) {
	book, author, err := processor.storage.GetBooksByAuthorId(id)

	if err != nil {
		return nil, processor.storage.NullAuthors(), errors.New("Author not found")
	}

	return book, author, nil

}

func (processor *AuthorsProcessor) FindAuthor(id int64) (*models.Authors, error) {
	author, err := processor.storage.GetAuthorById(id)

	if err != nil {
		return author, errors.New("user not found")
	}

	return author, nil

}

func (processor *AuthorsProcessor) StarTheAuthor(id int64) (*models.Authors, error) {
	err := processor.storage.PutStarByAuthorId(id)
	if err != nil {
		return processor.storage.NullAuthors(), errors.New("CANNOT DELETE BOOK")
	}

	fmt.Printf("Author %v has been stared.", id)
	author, _ := processor.FindAuthor(id)
	return author, nil
}

func (processor *AuthorsProcessor) DeleteAuthor(id int64) (int64, error) {
	deleted, err := processor.storage.DeleteAuthorById(id)
	if err != nil {
		return 0, errors.New("CANNOT DELETE THE AUTHOR")
	}
	fmt.Printf("Author %v has been deleted.", id)
	return deleted, nil
}

func (processor *AuthorsProcessor) UpdateAuthor(author *models.Authors) (*models.Authors, error) { //!!! ПРОВЕРИТЬ
	return processor.storage.ChangeAuthor(author)
}
