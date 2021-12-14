package main

import (
	"fmt"
	"sync"
)

//communicate by sharing memory (NOT ADVISED)

type OpCount struct {
	count int
	sync.Mutex
}

func (o *OpCount) Increment() {
	o.Lock()
	{
		o.count++
	}
	o.Unlock()
}

var opCount OpCount
var wg sync.WaitGroup

var mutex sync.Mutex

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go add(100, 200)
	}
	wg.Wait()
	fmt.Println(opCount.count)
}

func add(x, y int) {
	fmt.Println(x + y)
	opCount.Increment()
	wg.Done()
}
