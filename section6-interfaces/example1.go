package main

import "fmt"

type user struct {
	name string
}

func (u *user) print() {
	fmt.Printf("user{name: %s}\n", u.name)
}

type printable interface {
	print()
}

func printData(object printable) {
	object.print()
}

func main() {
	a := user{name: "Vatal"}
	printData(&a)
}
