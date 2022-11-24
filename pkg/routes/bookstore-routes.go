package routes

import (
	"github.com/furkannarin/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterbookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/", controllers.DeleteBook).Methods("DELETE")
}
