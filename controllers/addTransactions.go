package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jashfeer/transationAPIs/models"
)

var transactions = []models.Transaction{}
var amountInOrder = []models.Transaction{}
var status models.Status
var timesMap = make(map[time.Time]bool)
var amountMap = make(map[float64]bool)

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
	fmt.Println("UTC time now: ", timeNow)
	difference := timeNow.Sub(transaction.Timestamp)
	ds := difference.Seconds()

	if ds < 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	} else if ds > 60 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Set transaction time in ascending order
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

	// Set transaction amount in ascending order
	amountL := len(amountInOrder)
	if amountL == 0 {
		amountInOrder = append(amountInOrder, transaction)
	} else {
		flg := true
		for i, value := range amountInOrder {
			ok := transaction.Amount < value.Amount
			if ok {
				if i == 0 {
					amountInOrder = append([]models.Transaction{transaction}, amountInOrder...)
				} else {
					amountInOrder = append(amountInOrder, transaction)
					copy(amountInOrder[i:], amountInOrder[i-1:])
					amountInOrder[i] = transaction
				}
				flg = false
				break
			}
		}
		if flg {
			amountInOrder = append(amountInOrder, transaction)
		}
	}

	//Updating Sum,Count,timeMap and amountMap
	status.Sum += transaction.Amount
	status.Count++
	timesMap[transaction.Timestamp] = true
	amountMap[transaction.Amount] = true

	w.WriteHeader(http.StatusCreated)
}
