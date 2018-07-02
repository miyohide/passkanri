package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/miyohide/passkanri/cmd"
)

func main() {
	db, err := sql.Open("sqlite3", "db/passkanri.sqlite3")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "passkanri" ("id" INTEGER PRIMARY KEY AUTOINCREMENT, "name" VARCHAR(255), "url" VARCHAR(255), "password" VARCHAR(255))`,
	)
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
