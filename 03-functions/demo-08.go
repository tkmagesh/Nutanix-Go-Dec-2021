package main

import "fmt"

func main() {
	fmt.Println("main triggered")
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	defer f1()
	defer f2()
	defer f3()

	x, y := 100, 0
	result := x / y
	fmt.Println("result = ", result)
	fmt.Println("exiting from main")
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func f3() {
	fmt.Println("f3 invoked")
}
