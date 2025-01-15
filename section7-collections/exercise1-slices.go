package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:2]
	fmt.Println(slice, len(slice), cap(slice))
	slice = slice[:len(arr)-1]
	fmt.Println(slice, len(slice), cap(slice))
}
