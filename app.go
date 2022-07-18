package main

import (
	//"fmt"
	//"log"
	"database/sql"

	"github.com/gorilla/mux"
	//"github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {

}

func (a *App) Run(addr string) {

}
