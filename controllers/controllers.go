package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jashfeer/transationAPIs/models"
)

var transactions = []models.Transaction{}

//Adding transactions
func AddTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transaction models.Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(400)
		return
	}

	timeNow := time.Now().UTC()
	difference := timeNow.Sub(transaction.Timestamp)

	if difference.Seconds() < 0 {
		json.NewEncoder(w).Encode(422)
		return
	} else if difference.Seconds() > 60 {
		json.NewEncoder(w).Encode(204)
		return
	}

	transactions = append(transactions, transaction)
	json.NewEncoder(w).Encode(201)
}

//Geting statistics of transactions in last 60 seconds
func GetStatistics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("transactions:   ", transactions)
}

//Deleting all transactions
func DeleteTransactions(w http.ResponseWriter, r *http.Request) {

	transactions = transactions[:0]
	json.NewEncoder(w).Encode(204)

}
