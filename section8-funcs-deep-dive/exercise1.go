package main

import "fmt"

type number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func multiplyBy[T number](num T, multiplier T) T {
	return num * multiplier
}

func filterBy[T number](filter func(T) bool, arr ...T) []T {
	newArr := arr[0:0]
	for _, val := range arr {
		if filter(val) {
			newArr = append(newArr, val)
		}
	}
	return newArr
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	odds := filterBy(func(num int) bool {
		return num%2 != 0
	}, nums...)
	for _, value := range odds {
		fmt.Printf("doubled odd num of %d is %d\n", value, multiplyBy(value, 2))
	}
}
