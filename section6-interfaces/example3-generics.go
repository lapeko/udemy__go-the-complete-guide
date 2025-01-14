package main

import (
	"fmt"
	"reflect"
)

type number interface {
	int | int8 | int16 | int32 | int64 | float64 | float32
}

func add[T number](a T, b T) T {
	return a + b
}

func sum[T int | int8 | int16 | int32 | int64 | float64 | float32](a T, b T) T {
	return a + b
}

func main() {
	res := add(2, 3)
	fmt.Println("add(2, 3) is ", res, reflect.TypeOf(res).Name())
	res2 := sum(2.2, 3)
	fmt.Println("sum(2.2, 3) is ", res2, reflect.TypeOf(res2).Name())
}
