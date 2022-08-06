package handlers

import (
	"encoding/json"
	"net/http"

	middl "github.com/wertick01/dclib/cmd/web/middleware"
	"github.com/wertick01/dclib/internals/app/models"
	"github.com/wertick01/dclib/internals/app/processors"
)

type FavorietesHandler struct {
	processor *processors.FavorietesProcessor
}

func NewFavorietesHandler(processor *processors.FavorietesProcessor) *FavorietesHandler {
	handler := new(FavorietesHandler)
	handler.processor = processor
	return handler
}

func (handler *FavorietesHandler) AddFavorieteBook(w http.ResponseWriter, r *http.Request) {
	var favorieteBook *models.FavorieteBooks

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&favorieteBook)
	if err != nil {
		WrapError(w, err)
		return
	}
	_, err = handler.processor.AddFavorieteBook(favorieteBook)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   favorieteBook,
	}

	WrapOK(w, m)
}

func (handler *FavorietesHandler) AddFavorieteAuthor(w http.ResponseWriter, r *http.Request) {
	var favorieteAuthor *models.FavorieteAuthors

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&favorieteAuthor)
	if err != nil {
		WrapError(w, err)
		return
	}

	_, err = handler.processor.AddFavorieteAuthor(favorieteAuthor)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   favorieteAuthor,
	}

	WrapOK(w, m)
}

func (handler *FavorietesHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	var favorietes *models.FavorieteBooks

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&favorietes)
	if err != nil {
		WrapError(w, err)
		return
	}

	list, err := handler.processor.ListFavorieteBooks(favorietes.UserId)

	if err != nil {
		WrapError(w, err)
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   list,
	}

	WrapOK(w, m)
}

func (handler *FavorietesHandler) ListAuthors(w http.ResponseWriter, r *http.Request) {
	//vars := r.URL.Query() ЗАЧЕМ ОНО ТУТ НАДО
	var favorietes *models.FavorieteAuthors

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&favorietes)
	if err != nil {
		WrapError(w, err)
		return
	}
	list, err := handler.processor.ListFavorieteAuthors(favorietes.UserId)

	if err != nil {
		WrapError(w, err)
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   list,
	}

	WrapOK(w, m)
}

func (handler *FavorietesHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {

	var favorietes *models.FavorieteBooks

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&favorietes)
	if err != nil {
		WrapError(w, err)
		return
	}

	deletedbook, err := handler.processor.DeleteFromFavorieteBooks(favorietes)
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

func (handler *FavorietesHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {

	var favorietes *models.FavorieteAuthors

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&favorietes)
	if err != nil {
		WrapError(w, err)
		return
	}

	deletedauthor, err := handler.processor.DeleteFromFavorieteAuthors(favorietes)
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
