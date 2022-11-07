package controllers

import "net/http"

//Deleting all transactions
func DeleteTransactions(w http.ResponseWriter, r *http.Request) {

	//clear all values from transactions
	transactions = transactions[:0]

	//clear all values from amountInOrder
	amountInOrder = amountInOrder[:0]

	//clear all values from status
	status.Avg = 0
	status.Sum = 0
	status.Count = 0
	status.Max = 0
	status.Min = 0

	//clear all values from timeMap
	for k := range timesMap {
		delete(timesMap, k)
	}

	//clear all values from amountMap
	for k := range amountMap {
		delete(amountMap, k)
	}

	w.WriteHeader(http.StatusNoContent)
}
