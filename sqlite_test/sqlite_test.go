package mylib

import "testing"

func TestAll(t *testing.T) {
	const dbpath = "foo.db"

	db := InitDB(dbpath)
	defer db.Close()
}
