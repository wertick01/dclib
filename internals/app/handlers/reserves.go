package handlers

import (
	"encoding/json"
	"net/http"

	middl "github.com/wertick01/dclib/cmd/web/middleware"
	"github.com/wertick01/dclib/internals/app/models"
	"github.com/wertick01/dclib/internals/app/processors"
)

type ReservesHandler struct {
	processor *processors.ReserveProcessor
}

func NewReservesHandler(processor *processors.ReserveProcessor) *ReservesHandler {
	handler := new(ReservesHandler)
	handler.processor = processor
	return handler
}

func (handler *ReservesHandler) List(w http.ResponseWriter, r *http.Request) {

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	list, err := handler.processor.GetList()

	if err != nil {
		WrapError(w, r, err)
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   list,
	}

	WrapOK(w, m)
}

func (handler *ReservesHandler) Reserve(w http.ResponseWriter, r *http.Request) {
	var reserved *models.Booking

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&reserved)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	_, err = handler.processor.ReserveBook(reserved)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   reserved,
	}

	WrapOK(w, m)
}

func (handler *ReservesHandler) Return(w http.ResponseWriter, r *http.Request) {
	var returned *models.Booking

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&returned)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	book, err := handler.processor.ReturnBook(returned)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   book,
	}

	WrapOK(w, m)
}

func (handler *ReservesHandler) Confirm(w http.ResponseWriter, r *http.Request) {
	var returned *models.Booking

	w, r, err := middl.CheckToken(w, r)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&returned)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	book, err := handler.processor.ConfirmBook(returned)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   book,
	}

	WrapOK(w, m)
}
