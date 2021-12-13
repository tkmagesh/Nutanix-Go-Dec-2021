package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Deferred @ main - 1")
	}()
	defer func() {
		fmt.Println("Deferred @ main - 2")
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	defer func() {
		fmt.Println("Deferred @ f1 - 1")
	}()
	defer func() {
		fmt.Println("Deferred @ f1 - 2")
	}()
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer func() {
		fmt.Println("Deferred @ f2 - 1")
	}()
	defer func() {
		fmt.Println("Deferred @ f2 - 2")
	}()
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
