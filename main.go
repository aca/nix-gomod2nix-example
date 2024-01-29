package main

import (
	"database/sql"
	"log"
	_ "github.com/godror/godror"
)

func main() {
	var	dsn = `user="SYSTEM" password="password" charset=KO16MSWIN949 connectString="localhost:1521/FREEPDB1"`
	db, err := sql.Open("godror", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("done")
}

