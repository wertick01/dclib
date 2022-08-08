package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/wertick01/dclib/internals/app/auth"
	"github.com/wertick01/dclib/internals/app/models"
	"github.com/wertick01/dclib/internals/app/processors"
)

type UsersHandler struct {
	processor *processors.UsersProcessor
	authorise *auth.Authoriser
}

func NewUsersHandler(processor *processors.UsersProcessor) *UsersHandler {
	handler := new(UsersHandler)
	handler.processor = processor
	return handler
}

var mySigningKey = []byte("johenews")

func (handler *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser *models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	user, err := handler.processor.CreateUser(newUser)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	createduser := handler.CreateUserResponseHelper(user)

	var m = map[string]interface{}{
		"result": "OK",
		"data":   createduser,
	}

	expirationTime := time.Now().Add(1 * time.Minute)

	claims := &models.Claims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	WrapOK(w, m)
}

func (handler *UsersHandler) List(w http.ResponseWriter, r *http.Request) {

	//w, r = CheckToken(w, r)

	list, err := handler.processor.ListUsers()

	if err != nil {
		WrapError(w, r, err)
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   list,
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		WrapError(w, r, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	user, err := handler.processor.FindUser(id)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   user,
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) FindByPhone(w http.ResponseWriter, r *http.Request) {
	var wantedUser *models.User

	err := json.NewDecoder(r.Body).Decode(&wantedUser)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	user, err := handler.processor.FindByPhone(wantedUser.Phone)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   user,
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) Change(w http.ResponseWriter, r *http.Request) {
	var changeUser *models.User

	err := json.NewDecoder(r.Body).Decode(&changeUser)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	user, err := handler.processor.UpdateUser(changeUser)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   user,
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		WrapError(w, r, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	deleteduser, err := handler.processor.DeleteUser(id)
	if err != nil {
		WrapError(w, r, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   deleteduser,
	}

	WrapOK(w, m)
}
