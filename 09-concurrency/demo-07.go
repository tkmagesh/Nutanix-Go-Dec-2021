package main

import (
	"fmt"
)

func main() {
	resultCh := make(chan int)
	go add(100, 200, resultCh)
	result := <-resultCh
	fmt.Println("Result:", result)
}

func add(x, y int, resultCh chan int) {
	fmt.Printf("Processing %d & %d\n", x, y)
	result := x + y
	resultCh <- result
}
