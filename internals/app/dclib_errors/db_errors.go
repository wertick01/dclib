package dclib_errors

import (
	"errors"
)

type AuthorError struct {
	author_error error
}

var Author_error_create error = errors.New("#1 {ERROR_WITH_CREATING_AUTHOR} --> Что-то введено не правильно или возникла внутренняя ошибка.")
var Author_error_list error = errors.New("#2 {ERROR_WITH_GETTING_AUTHORS_LIST} --> Невозможно получить список авторов.")
var Author_error_getbyid error = errors.New("#3 {ERROR_WITH_GETTING_AUTHOR_BY_ID} --> Невозможно найти автора.")
var Author_error_getauthbooks error = errors.New("#4 {ERROR_WITH_GETTING_BOOKS_BY_AUTHOR_ID} --> Невозможно получить книги автора.")

//var Author_error_cantfind error = errors.New("#5 {ERROR_WITH_GETTING_BOOKS_BY_AUTHOR_ID} --> Невозможно получить книги автора.")

/*
func (autherr *AuthorError) DBErrorConstructor(code, ticker, text string) error {
	var author_err = &models.ErrorModel{}
	return author_err{author_error: errors.New(fmt.Sprintf("%v {%v} --> %v", code, ticker, text))}
}
*/
