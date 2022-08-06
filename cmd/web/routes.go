package api

import (
	"github.com/gorilla/mux"
	"github.com/wertick01/dclib/internals/app/auth"
	"github.com/wertick01/dclib/internals/app/handlers"
)

func CreateRoutes(
	booksHandler *handlers.BooksHandler,
	usersHandler *handlers.UsersHandler,
	authorsHandler *handlers.AuthorsHandler,
	authoriser *auth.Authoriser,
	favorietesHandler *handlers.FavorietesHandler,
	reseervesHandler *handlers.ReservesHandler,
) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/api/login", authoriser.Login).Methods("POST")
	r.HandleFunc("/api/refresh", authoriser.Refresh).Methods("GET")
	r.HandleFunc("/api/registration", usersHandler.Create).Methods("POST")

	r.HandleFunc("/api/users", usersHandler.List).Methods("GET")
	r.HandleFunc("/api/users", usersHandler.Change).Methods("PUT")
	r.HandleFunc("/api/users/phone", usersHandler.FindByPhone).Methods("POST")
	r.HandleFunc("/api/users/{id:[0-9]+}", usersHandler.Find).Methods("GET")
	r.HandleFunc("/api/users/{id:[0-9]+}", usersHandler.Delete).Methods("DELETE")
	//r.HandleFunc("/api/send-code").Methods("POST")
	//r.HandleFunc("/api/refresh-token").Methods("POST")
	//r.HandleFunc("/api/upload").Methods("POST")
	//r.HandleFunc("/api/img/upload").Methods("POST")
	//r.HandleFunc("/api/img/destroy").Methods("POST")
	//r.HandleFunc("/api/whoami").Methods("GET")

	r.HandleFunc("/api/book", booksHandler.Create).Methods("POST")
	r.HandleFunc("/api/book", booksHandler.List).Methods("GET")
	r.HandleFunc("/api/book", booksHandler.Change).Methods("PUT")
	r.HandleFunc("/api/book/{id:[0-9]+}", booksHandler.Find).Methods("GET")
	//r.HandleFunc("/api/book/{id:[0-9]+}", booksHandler.Star).Methods("POST")
	r.HandleFunc("/api/book/{id:[0-9]+}", booksHandler.Delete).Methods("DELETE")

	r.HandleFunc("/api/authors", authorsHandler.List).Methods("GET")
	r.HandleFunc("/api/authors", authorsHandler.Create).Methods("POST")
	r.HandleFunc("/api/authors", authorsHandler.Change).Methods("PUT")
	r.HandleFunc("/api/authors/{id:[0-9]+}", authorsHandler.Find).Methods("GET")
	//r.HandleFunc("/api/authors/{id:[0-9]+}", authorsHandler.Star).Methods("POST")
	r.HandleFunc("/api/authors/{id:[0-9]+}", authorsHandler.Delete).Methods("DELETE")
	r.HandleFunc("/api/authors/books/{id:[0-9]+}", authorsHandler.FindBooks).Methods("GET")

	r.HandleFunc("/api/favorietes/books/list", favorietesHandler.ListBooks).Methods("POST")
	r.HandleFunc("/api/favorietes/authors/list", favorietesHandler.ListAuthors).Methods("POST")
	r.HandleFunc("/api/favorietes/books/add", favorietesHandler.AddFavorieteBook).Methods("POST")
	r.HandleFunc("/api/favorietes/authors/add", favorietesHandler.AddFavorieteAuthor).Methods("POST")
	r.HandleFunc("/api/favorietes/books/delete", favorietesHandler.DeleteBook).Methods("POST")
	r.HandleFunc("/api/favorietes/authors/delete", favorietesHandler.DeleteAuthor).Methods("POST")

	r.HandleFunc("/api/reserved/list", reseervesHandler.List).Methods("GET")
	r.HandleFunc("/api/reserved/reserve", reseervesHandler.Reserve).Methods("POST")
	r.HandleFunc("/api/reserved/return", reseervesHandler.Return).Methods("PUT")
	r.HandleFunc("/api/reserved/confirm", reseervesHandler.Confirm).Methods("PUT")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler() //оборачиваем 404, для обработки NotFound
	return r
}
