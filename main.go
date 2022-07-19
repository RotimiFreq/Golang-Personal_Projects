package main

import (
	"database/sql"
	"log"
)

const (
	db_user     = "postgres"
	db_password = "ethanol"
	db_name     = "test"
)

func checkErr(err error) {
	if err != nil {
		log.Printf("the error is :", err)

	}

}
func initDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://samuel:sam123@localhost/soundAPI_DB?sslmode=disable")
	checkErr(err)

	stmt, err := db.Prepare("CREATE TABLE PRODUCTINFO(ID SERIAL PRIMARY KEY, NAME TEXT NOT NULL);")

	res, err := stmt.Exec()

	log.Println(res)
	checkErr(err)

	return db, nil

}

func main() {
	initDB()
}
