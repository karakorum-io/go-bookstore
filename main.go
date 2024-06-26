package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karakorum.io/karakorum-book-store/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
