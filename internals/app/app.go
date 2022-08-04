package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	api "github.com/wertick01/dclib/cmd/web"
	"github.com/wertick01/dclib/cmd/web/middleware"
	"github.com/wertick01/dclib/internals/app/auth"
	db3 "github.com/wertick01/dclib/internals/app/db"
	"github.com/wertick01/dclib/internals/app/handlers"
	"github.com/wertick01/dclib/internals/app/processors"
	"github.com/wertick01/dclib/internals/cfg"
)

type AppServer struct {
	config cfg.Cfg
	srv    *http.Server
	db     *sql.DB
	ctx    context.Context
}

func NewServer(config cfg.Cfg, ctx context.Context) *AppServer { //задаем поля нашего сервера, для его старта нам нужен контекст и конфигурация
	server := new(AppServer)
	server.ctx = ctx
	server.config = config
	return server
}

func (server *AppServer) Serve() {
	log.Println("Starting server")
	log.Println(server.config.GetDBString())
	var err error
	server.db, err = openDB(server.config.GetDBString())
	if err != nil {
		log.Fatalln(err)
	}
	//defer server.db.Close()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	booksStorage := db3.NewBooksStorage(server.db)
	usersStorage := db3.NewUsersStorage(server.db)
	authorsStorage := db3.NewAuthorsStorage(server.db)
	favorietesStorage := db3.NewFavorietesStorage(server.db)
	reserveStorage := db3.NewBooksStorage(server.db)

	booksProcessor := processors.NewBooksProcessor(booksStorage)
	usersProcessor := processors.NewUsersProcessor(usersStorage)
	authorsProcessor := processors.NewAuthorsProcessor(authorsStorage)
	favorietesProcessor := processors.NewFavorietesProcessor(favorietesStorage)
	reserveProcessor := processors.NewReserveProcessor(reserveStorage)

	booksHandler := handlers.NewBooksHandler(booksProcessor)
	usersHandler := handlers.NewUsersHandler(usersProcessor)
	authorsHandler := handlers.NewAuthorsHandler(authorsProcessor)
	authoriser := auth.NewAuthoriser(usersProcessor)
	favorietesHandler := handlers.NewFavorietesHandler(favorietesProcessor)
	reserveHandler := handlers.NewReservesHandler(reserveProcessor)

	routes := api.CreateRoutes(
		booksHandler,
		usersHandler,
		authorsHandler,
		authoriser,
		favorietesHandler,
		reserveHandler,
	) //хендлеры напрямую используются в путях
	routes.Use(middleware.RequestLog) //middleware используем здесь, хотя можно было бы и в CreateRoutes

	server.srv = &http.Server{ //в отличие от примеров http, здесь мы передаем наш server в поле структуры, для работы в Shutdown
		//Addr: ":8080",
		Addr:     ":" + server.config.Port,
		Handler:  routes,
		ErrorLog: errorLog,
	}

	log.Println("Server started")
	err = server.srv.ListenAndServe() //запускаем сервер

	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func (server *AppServer) Shutdown() {
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.db.Close() //закрываем соединение с БД
	defer func() {
		cancel()
	}()
	var err error
	if err = server.srv.Shutdown(ctxShutDown); err != nil { //выключаем сервер, с ограниченным по времени контекстом
		log.Fatalf("server Shutdown Failed:%s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
