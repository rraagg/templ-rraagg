package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/a-h/templ"
)

type user struct {
	firstName []byte
	lastName  []byte
	email     []byte
}

func (u user) FirstName() []byte {
	return u.firstName
}

func (u user) LastName() []byte {
	return u.lastName
}

func (u user) FullName() []byte {
	sliceOfByteSlices := [][]byte{u.firstName, u.lastName}
	return bytes.Join(sliceOfByteSlices, []byte(" "))
}

func (u user) Email() []byte {
	return u.email
}

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStatement := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name TEXT
	);
	`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table users created successfully")
	webUser := user{
		firstName: []byte("rraagg"),
		lastName:  []byte("thegreat"),
	}
	component := hello(string(webUser.FullName()))
	http.Handle("/", templ.Handler(component))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
