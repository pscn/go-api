package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pscn/flavor2go/api"
	"github.com/pscn/flavor2go/middleware"
	"github.com/pscn/flavor2go/model"

	"github.com/karrick/golf"
)

var port = golf.IntP('l', "listen", 3000, "port to listen on")

// FIXME: maybe just use a connect string
// FIXME:² read from config
var pgHost = golf.StringP('h', "host", "localhost", "postgres host")
var pgPort = golf.IntP('p', "port", 5432, "postgres port")
var pgDatabase = golf.StringP('d', "database", "gusta", "postgres database")
var pgUser = golf.StringP('u', "user", "gusta", "postgres user")
var pgPass = golf.String("pass", "changeme", "postgres password")
var pgUseSSL = golf.BoolP('s', "ssl", false, "postgres use SSL")

func main() {
	golf.Parse()
	log.SetFlags(0)

	m := model.NewPostgres(*pgHost, *pgPort, *pgUser, *pgDatabase, *pgPass, *pgUseSSL)
	a := api.New(m)
	defer m.Close()

	r := mux.NewRouter()
	a.Register(r)

	r.Use(middleware.Log)
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("listen on: %d", *port)
	log.Fatal(http.ListenAndServe(addr, r))
}
