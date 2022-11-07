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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	timeNow := time.Now().UTC()
	fmt.Println("NOW: ", timeNow)
	difference := timeNow.Sub(transaction.Timestamp)

	if difference.Seconds() < 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	} else if difference.Seconds() > 60 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	l := len(transactions)
	if l == 0 {
		transactions = append(transactions, transaction)
	} else {
		flg := true
		for i, value := range transactions {
			ok := transaction.Timestamp.Before(value.Timestamp)
			if ok {
				if i == 0 {
					transactions = append([]models.Transaction{transaction}, transactions...)
				} else {
					transactions = append(transactions, transaction)
					copy(transactions[i:], transactions[i-1:])
					transactions[i] = transaction
				}
				flg = false
				break
			}
		}
		if flg {
			transactions = append(transactions, transaction)
		}
	}

	w.WriteHeader(http.StatusCreated)
}

//Geting statistics of transactions in last 60 seconds
func GetStatistics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("transactions:   ", transactions)
}

//Deleting all transactions
func DeleteTransactions(w http.ResponseWriter, r *http.Request) {

	transactions = transactions[:0]
	w.WriteHeader(http.StatusNoContent)
}
