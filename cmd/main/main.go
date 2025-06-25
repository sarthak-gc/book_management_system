package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthak-gc/book-management-system/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.BookRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe(":9090", r)
}
