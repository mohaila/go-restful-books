package main

import (
	"database/sql"
)

type bookStore interface {
	migrate() (e error)
	all() (books []Book, e error)
	create(b Book) (id int64, e error)
	get(id int64) (b Book, e error)
	update(b Book) (e error)
	delete(id int64) (e error)
}

type bookSQLStore struct {
	db *sql.DB
}

func newBookSQLStore(db *sql.DB) (store bookStore) {
	store = &bookSQLStore{db: db}
	return
}

func (store *bookSQLStore) migrate() (e error) {
	// Create table
	st, e := store.db.Prepare(`
		create table if not exists books (
		id integer primary key, 
		title varchar(64),
		author varchar(64)
	)`)
	if e != nil {
		return
	}
	_, e = st.Exec()

	return
}

func (store *bookSQLStore) all() (books []Book, e error) {
	rows, e := store.db.Query(`
		select * from books
	`)
	if e != nil {
		return
	}
	defer rows.Close()

	books = []Book{}
	var book Book

	for rows.Next() {
		rows.Scan(&book.ID, &book.Title, &book.Author)
		books = append(books, book)
	}

	return
}

func (store *bookSQLStore) create(b Book) (id int64, e error) {
	st, e := store.db.Prepare(`
		insert into books (title, 
			author)
		values (?,
			?
		)
	`)
	if e != nil {
		return
	}
	result, e := st.Exec(b.Title, b.Author)
	if e != nil {
		return
	}

	id, e = result.LastInsertId()
	return
}

func (store *bookSQLStore) get(id int64) (b Book, e error) {
	e = store.db.QueryRow(`
		select * from books 
		where id = ?
	`, id).Scan(&b.ID, &b.Title, &b.Author)

	return
}

func (store *bookSQLStore) update(b Book) (e error) {
	st, e := store.db.Prepare(`
		update books
		set title = ?,
			author = ?
		where id = ?
	`)
	if e != nil {
		return
	}

	_, e = st.Exec(b.Title, b.Author, b.ID)
	return
}

func (store *bookSQLStore) delete(id int64) (e error) {
	st, e := store.db.Prepare(`
		delete from books
		where id = ?
	`)

	if e != nil {
		return
	}

	_, e = st.Exec(id)

	return
}
