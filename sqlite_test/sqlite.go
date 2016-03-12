package mylib

import (
	"database/sql"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type TestItem struct {
	Id    string
	Name  string
	Phone string
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}
