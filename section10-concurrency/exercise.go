package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func printNum(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("printNum: ", num)
}

func printName(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("printName: ", name)
}

func heavySumCalculation(num1 int, num2 int, done chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	done <- num1 + num2
}

func main() {
	fmt.Println("main start")
	wg := &sync.WaitGroup{}
	sumChan := make(chan int, 1)

	wg.Add(1)
	go printNum(rand.Intn(100), wg)

	wg.Add(1)
	go printName("Me", wg)

	a, b := rand.Intn(100), rand.Intn(100)
	wg.Add(1)
	go heavySumCalculation(a, b, sumChan, wg)

	wg.Wait()
	sum := <-sumChan

	fmt.Printf("heavySumCalculation of %d and %d is %d\n", a, b, sum)
	fmt.Println("main finish")
}
