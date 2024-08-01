package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getRequest).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}

func getRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "get\n")
}

func postRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "post\n")
}
