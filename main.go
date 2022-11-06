package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jashfeer/transationAPIs/controllers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/transactions", controllers.AddTransactions).Methods("POST")
	r.HandleFunc("/statistics", controllers.GetStatistics).Methods("GET")
	r.HandleFunc("/transactions", controllers.DeleteTransactions).Methods("DELETE")

	log.Println("Runing...in port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
