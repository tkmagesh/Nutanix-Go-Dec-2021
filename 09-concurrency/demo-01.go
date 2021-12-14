package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid())
	fmt.Println("main started")
	go f1() //=> goroutine
	f2()

	//DO NOT DO THIS
	//time.Sleep(time.Second)

	var input string
	fmt.Scanln(&input)
	fmt.Println("main completed")
}

func f1() {

	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
