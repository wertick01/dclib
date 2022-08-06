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

type AuthorsHandler struct {
	processor *processors.AuthorsProcessor
}

func NewAuthorsHandler(processor *processors.AuthorsProcessor) *AuthorsHandler {
	handler := new(AuthorsHandler)
	handler.processor = processor
	return handler
}

func (handler *AuthorsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newAuthor *models.Authors

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&newAuthor)
	if err != nil {
		WrapError(w, err)
		return
	}

	author, err := handler.processor.CreateAuthor(newAuthor)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   author,
	}

	WrapOK(w, m)
}

func (handler *AuthorsHandler) List(w http.ResponseWriter, r *http.Request) {

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	list, err := handler.processor.ListAuthors()

	if err != nil {
		WrapError(w, err)
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   list,
	}

	WrapOK(w, m)
}

func (handler *AuthorsHandler) Find(w http.ResponseWriter, r *http.Request) {

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	vars := mux.Vars(r) //переменные, обьявленные в ресурсах попадают в Vars и могут быть адресованы
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	author, err := handler.processor.FindAuthor(id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   author,
	}

	WrapOK(w, m)
}

func (handler *AuthorsHandler) FindBooks(w http.ResponseWriter, r *http.Request) {

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}
	//go middl.TimeChecker(w, r)

	vars := mux.Vars(r) //переменные, обьявленные в ресурсах попадают в Vars и могут быть адресованы
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	books, author, err := handler.processor.AuthorsBooks(id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"books":  books,
		"author": author,
	}

	WrapOK(w, m)
}

func (handler *AuthorsHandler) Change(w http.ResponseWriter, r *http.Request) {

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	var author *models.Authors

	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		WrapError(w, err)
		return
	}

	updatedauthor, err := handler.processor.UpdateAuthor(author)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   updatedauthor,
	}

	WrapOK(w, m)
}

func (handler *AuthorsHandler) Star(w http.ResponseWriter, r *http.Request) {

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

	author, err := handler.processor.StarTheAuthor(id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"author": author,
	}

	WrapOK(w, m)
}

func (handler *AuthorsHandler) Delete(w http.ResponseWriter, r *http.Request) {

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	vars := mux.Vars(r) //переменные, обьявленные в ресурсах попадают в Vars и могут быть адресованы
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	deletedauthor, err := handler.processor.DeleteAuthor(id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   deletedauthor,
	}

	WrapOK(w, m)
}
