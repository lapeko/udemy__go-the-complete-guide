package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("", reflect.TypeOf(arr))
	slice := arr[1:2]
	fmt.Printf("Init slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
	slice = slice[:len(arr)-1]
	fmt.Printf("Extended slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
}
