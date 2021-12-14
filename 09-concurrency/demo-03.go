package main

import (
	"fmt"
	"sync"
)

//communicate by sharing memory (NOT ADVISED)

var opCount int
var wg sync.WaitGroup

var mutex sync.Mutex

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go add(100, 200)
	}
	wg.Wait()
	fmt.Println(opCount)
}

func add(x, y int) {
	fmt.Println(x + y)
	mutex.Lock()
	{
		opCount++
	}
	mutex.Unlock()
	wg.Done()
}
