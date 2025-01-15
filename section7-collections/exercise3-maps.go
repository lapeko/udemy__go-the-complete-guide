package main

import "fmt"

type distanceMap map[string]float64

func (d distanceMap) output() {
	for key, value := range d {
		fmt.Printf("%s: %.3f\n", key, value)
	}
}

func main() {
	distancesMap := make(distanceMap, 3)
	distancesMap["Portugal"] = 3202.22
	distancesMap["Spain"] = 2635.16
	distancesMap["Greece"] = 1248.39
	distancesMap.output()
}
