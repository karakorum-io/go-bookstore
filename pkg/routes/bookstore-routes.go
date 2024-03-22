package routes

import (
	"github.com/gorilla/mux"
	"github.com/karakorum.io/karakorum-book-store/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
