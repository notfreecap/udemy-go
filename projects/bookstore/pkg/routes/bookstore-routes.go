package routes

import (
	"github.com/gorilla/mux"
	"github.com/notfreecap/bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PATCH")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
