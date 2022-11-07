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

	for _, value := range transactions {
		difference := timeNow.Sub(value.Timestamp)
		ds := difference.Seconds()
		if ds < 60 {
			break
		} else {
			//Updating Sum,Count,timeMap and amountMap
			status.Sum -= value.Amount
			status.Count--
			timesMap[value.Timestamp] = false
			amountMap[value.Amount] = false
		}
	}

	//updating avg
	status.Avg = status.Sum / float64(status.Count)

	min := true
	max := true
	for i, j := 0, len(amountInOrder)-1; i <= j; i, j = i+1, j-1 {
		minValue_1 := timesMap[amountInOrder[i].Timestamp]
		minValue_2 := amountMap[amountInOrder[i].Amount]
		maxValue_1 := timesMap[amountInOrder[j].Timestamp]
		maxValue_2 := amountMap[amountInOrder[j].Amount]

		if minValue_1 && minValue_2 && min {
			status.Min = amountInOrder[i].Amount
			min = false
		}
		if maxValue_1 && maxValue_2 && max {
			status.Max = amountInOrder[j].Amount
			max = false
		}
		if !min && !max {
			break
		}
	}

	json.NewEncoder(w).Encode(&status)
}
