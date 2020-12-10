package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

// Connect to a mysql database in localhost:3306
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/task_database")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	// CreateDatabase(db) // uncomment to create database and reset table
	con = db
	return db
}

// função que recebe o banco de dados e uma query string, retorna resultado
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}
