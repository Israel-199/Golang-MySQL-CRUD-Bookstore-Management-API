package main

import (
	"fmt"
	"log"
	"net/http"

	"bookstore/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server running at the port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
