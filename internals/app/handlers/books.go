package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	middl "github.com/wertick01/dclib/cmd/web/middleware"
	"github.com/wertick01/dclib/internals/app/models"
	"github.com/wertick01/dclib/internals/app/processors"
)

type BooksHandler struct {
	processor *processors.BooksProcessor
}

func NewBooksHandler(processor *processors.BooksProcessor) *BooksHandler {
	handler := new(BooksHandler)
	handler.processor = processor
	return handler
}

func (handler *BooksHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newBook *models.Books

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		WrapError(w, err)
		return
	}

	_, err = handler.processor.CreateBook(newBook)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   "",
	}

	WrapOK(w, m)
}

func (handler *BooksHandler) List(w http.ResponseWriter, r *http.Request) {
	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	list, err := handler.processor.ListBooks()

	if err != nil {
		WrapError(w, err)
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   list,
	}

	WrapOK(w, m)
}

func (handler *BooksHandler) Find(w http.ResponseWriter, r *http.Request) {
	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	vars := mux.Vars(r)
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	book, err := handler.processor.FindBook(id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   book,
	}

	WrapOK(w, m)
}

func (handler *BooksHandler) Change(w http.ResponseWriter, r *http.Request) {
	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	var updateBook *models.Books

	err = json.NewDecoder(r.Body).Decode(&updateBook)
	if err != nil {
		WrapError(w, err)
		return
	}

	book, err := handler.processor.UpdateBook(updateBook)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   book,
	}

	WrapOK(w, m)
}

/*
func (handler *BooksHandler) Star(w http.ResponseWriter, r *http.Request) {

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	vars := mux.Vars(r)
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	book, err := handler.processor.StarTheBook(id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"author": book,
	}

	WrapOK(w, m)
}
*/

func (handler *BooksHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	vars := mux.Vars(r)
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	deletedbook, err := handler.processor.DeleteBook(id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   deletedbook,
	}

	WrapOK(w, m)
}
