package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Deferred @ main - 1") //=>12
	}()
	defer func() {
		fmt.Println("Deferred @ main - 2") //=> 11
	}()
	fmt.Println("main started") //=> 1
	f1()
	fmt.Println("main completed") //=> 10
}

func f1() {
	defer func() {
		fmt.Println("Deferred @ f1 - 1") //=> 9
	}()
	defer func() {
		fmt.Println("Deferred @ f1 - 2") //=> 8
	}()
	fmt.Println("f1 started") //=> 2
	f2()
	fmt.Println("f1 completed") //=>7
}

func f2() {
	defer func() {
		fmt.Println("Deferred @ f2 - 1") //=>6
	}()
	defer func() {
		fmt.Println("Deferred @ f2 - 2") //=> 5
	}()
	fmt.Println("f2 started")   //=> 3
	fmt.Println("f2 completed") //=> 4
}
