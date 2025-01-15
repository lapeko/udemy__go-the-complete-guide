package main

import "fmt"

func main() {
	hobbies := [...]string{"swimming", "playing guitar", "learning languages"}
	fmt.Printf("First hobbie: %s. Rest: %v\n", hobbies[0], hobbies[1:])
	firstTwo := make([]string, 2)
	copy(firstTwo, hobbies[:2])
	fmt.Println("Created with make: ", firstTwo)
	firstTwo2 := hobbies[:2]
	fmt.Println("Created without make: ", firstTwo2)
	lastTwo := append(firstTwo[1:], hobbies[2])
	fmt.Println("Last two hobbies: ", lastTwo)
	lastTwo2 := firstTwo2[1:3]
	fmt.Println("Last two hobbies: ", lastTwo2)
	goals := []string{"Buy beer", "Drink beer"}
	goals[1] = "Buy more beer"
	goals = append(goals, "Drink all beer")
	fmt.Println("Your goals: ", goals)

	type Product struct {
		title string
		id    int
		price float64
	}
	products := []Product{
		{title: "Heineken", id: 1, price: 3.49},
		{title: "Corona Extra", id: 2, price: 3.69},
	}
	products = append(products, Product{"Zubr", 3, 2.99})
	fmt.Println(products)
}
