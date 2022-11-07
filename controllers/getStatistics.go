package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Geting statistics of transactions in last 60 seconds
func GetStatistics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("transactions:   ", transactions)
	fmt.Println()
	fmt.Println()
	fmt.Println("amount in order:   ", amountInOrder)
	fmt.Println()
	fmt.Println()
	fmt.Println("time MAP: ", timesMap)
	fmt.Println()
	fmt.Println()
	fmt.Println("Amount MAp: ", amountMap)
	fmt.Println()
	fmt.Println()

	w.Header().Set("Content-Type", "application/json")

	timeNow := time.Now().UTC()
	fmt.Println("UTC time now: ", timeNow)

	if len(transactions) == 0 || status.Count <= 0 {
		status.Avg = 0
		status.Sum = 0
		status.Count = 0
		status.Max = 0
		status.Min = 0
		json.NewEncoder(w).Encode(&status)
		return
	}

	for i, value := range transactions {
		difference := timeNow.Sub(value.Timestamp)
		ds := difference.Seconds()
		if ds < 60 {
			transactions = transactions[i:]
			break
		}
		if status.Count <= 0 {
			transactions = transactions[:0]

		}
		//Updating Sum,Count,timeMap and amountMap
		status.Sum -= value.Amount
		status.Count--
		delete(timesMap, value.Timestamp)
		delete(amountMap, value.Amount)

	}

	//updating avg
	status.Avg = status.Sum / float64(status.Count)

	for i := 0; i < len(amountInOrder); i++ {
		_, minValue_1 := timesMap[amountInOrder[i].Timestamp]
		_, minValue_2 := amountMap[amountInOrder[i].Amount]

		if minValue_1 && minValue_2 {
			status.Min = amountInOrder[i].Amount
			amountInOrder = amountInOrder[i:]
			break
		}
	}

	for j := len(amountInOrder) - 1; j >= 0; j-- {
		_, minValue_1 := timesMap[amountInOrder[j].Timestamp]
		_, minValue_2 := amountMap[amountInOrder[j].Amount]

		if minValue_1 && minValue_2 {
			status.Max = amountInOrder[j].Amount
			amountInOrder = amountInOrder[:j+1]
			break
		}
	}

	json.NewEncoder(w).Encode(&status)
}
