package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(2)
	go f1(wg) //=> goroutine
	f2()

	wg.Wait()
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {

	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
