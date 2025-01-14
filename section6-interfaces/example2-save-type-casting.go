package main

import (
	"fmt"
	"reflect"
)

func addTwoAndPrintIfNum(data interface{}) {
	fmt.Println("func addTwoAndPrintIfNum")
	if num, ok := data.(int); ok {
		fmt.Println(num + 2)
	} else {
		fmt.Println("Not a number")
	}
}

func printType(data interface{}) {
	fmt.Printf("func printType\nIs int: %t\n", reflect.TypeOf(data).Kind() == reflect.Int)
	fmt.Print("The type is ")
	switch data.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("neither int nor string")
	}
}

func main() {
	addTwoAndPrintIfNum(struct{ test int }{test: 123})
	addTwoAndPrintIfNum(10)
	printType(10)
}
