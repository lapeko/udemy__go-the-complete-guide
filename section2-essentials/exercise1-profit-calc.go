package main

import "fmt"

func main() {
	data := map[string]float64{
		"revenue":  -1,
		"expenses": -1,
		"taxRate":  -1,
	}

	for key := range data {
		fmt.Printf("Please enter %s: ", key)
		for data[key] < 0 {
			var input float64
			if _, err := fmt.Scan(&input); err != nil {
				fmt.Printf("\nYou entered incorrect value. Please try againg to enter %s: ", key)
				continue
			}
			data[key] = input
		}
	}

	ebt := data["revenue"] - data["expenses"]
	profit := ebt * (1 - data["taxRate"]/100)
	ratio := ebt / profit
	fmt.Printf("\nebt: %.2f profit: %.2f, ratio: %.2f\n", ebt, profit, ratio)
}
