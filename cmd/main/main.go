package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/furkannarin/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterbookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
