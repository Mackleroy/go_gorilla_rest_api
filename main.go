package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/api"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/categories", api.GetCategoriesHandler).Methods("GET")
	r.HandleFunc("/categories", api.CreateCategoryHandler).Methods("POST")
	r.HandleFunc("/goods", api.CreateGoodHandler).Methods("POST")
	log.Println("Server is running")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
		return
	}

}
