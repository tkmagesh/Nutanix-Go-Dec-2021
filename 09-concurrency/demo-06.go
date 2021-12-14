package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	resultCh := make(chan int)
	go add(100, 200, wg, resultCh)
	result := <-resultCh
	wg.Wait()
	fmt.Println("Result:", result)
}

func add(x, y int, wg *sync.WaitGroup, resultCh chan int) {
	fmt.Printf("Processing %d & %d\n", x, y)
	result := x + y
	resultCh <- result
	wg.Done()
}
