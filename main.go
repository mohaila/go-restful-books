package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, e := sql.Open("sqlite3", "books.sqlite")
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()

	store := newBookSQLStore(db)
	e = store.migrate()
	if e != nil {
		log.Fatal(e)
	}

	container := restful.NewContainer()

	handler := newBookHandler(store)
	handler.regsiter(container)

	log.Printf("start listening on localhost:8000")
	server := &http.Server{
		Addr:    ":8000",
		Handler: container,
	}
	log.Fatal(server.ListenAndServe())

}
