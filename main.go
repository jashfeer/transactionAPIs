package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jashfeer/transationAPIs/controllers"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helooo")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")

	r.HandleFunc("/transactions", controllers.AddTransactions).Methods("POST")
	r.HandleFunc("/statistics", controllers.GetStatistics).Methods("GET")
	r.HandleFunc("/transactions", controllers.DeleteTransactions).Methods("DELETE")

	log.Println("Runing...in port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
