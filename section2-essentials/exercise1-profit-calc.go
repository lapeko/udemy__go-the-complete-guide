package main

import (
	"fmt"
	"os"
)

func main() {
	data := map[string]float64{
		"revenue":  -1,
		"expenses": -1,
		"taxRate":  -1,
	}

	for key := range data {
		fmt.Printf("Please enter %s: ", key)
		var input float64
		for {
			errMsg := fmt.Sprintf("You entered an incorrect value for %s. Please provide a positive number: ", key)
			if _, err := fmt.Scan(&input); err == nil && input > 0 {
				data[key] = input
				break
			}
			fmt.Print(errMsg)
		}
	}

	ebt := data["revenue"] - data["expenses"]
	profit := ebt * (1 - data["taxRate"]/100)
	ratio := ebt / profit
	stringOutput := fmt.Sprintf("ebt: %.2f profit: %.2f, ratio: %.2f\n", ebt, profit, ratio)
	fmt.Println(stringOutput)
	_ = os.WriteFile("output.txt", []byte(stringOutput), 0644)
}
