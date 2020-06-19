package main

import (
	"database/sql"
)

var (
	db *sql.DB
)

func checkDBError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:lqdLQD!!@(127.0.0.1:3306)/Game?charset=utf8")
	checkDBError(err)
}

func checkIdAndPassword(id, password string) bool {
	var rows, err = db.Query("select 1 from User where id=? and password=? limit 1", id, password)
	checkDBError(err)
	if rows.Next() {
		return true
	} else {
		return false
	}
}
