package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", cityhome).Methods("GET")
	r.HandleFunc("/cities", citieshandler).Methods("GET")
	r.HandleFunc("/populations", populationhandler).Methods("GET")

	http.ListenAndServe(":8185", r)
}

func cityhome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("cityhome")
	fmt.Fprintf(w, "cityhome : %s", r.URL.Path)
}

func citieshandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("citieshandler")
	fmt.Fprintf(w, "citieshandler : %s", r.URL.Path)
}

func populationhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("populationhandler")
	fmt.Fprintf(w, "populationhandler : %s", r.URL.Path)
}

func tmp{}
