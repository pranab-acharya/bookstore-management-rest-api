package main

import (
	"bookstore-management-rest-api/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	fmt.Println("Server is running on 9010")
	log.Fatal(http.ListenAndServe(":9010", r))
}
