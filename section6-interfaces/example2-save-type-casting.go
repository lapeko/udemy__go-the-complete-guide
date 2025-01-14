package main

import "fmt"

func addTwoAndPrintIfNum(data interface{}) {
	if num, ok := data.(int); ok {
		fmt.Println(num + 2)
	} else {
		fmt.Println("Not a number")
	}
}

func main() {
	addTwoAndPrintIfNum(struct{ test int }{test: 123})
	addTwoAndPrintIfNum(10)
}
