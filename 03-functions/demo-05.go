package main

import "fmt"

func main() {
	/*
		fmt.Println(add(100, 200))
		fmt.Println(subtract(100, 200))
	*/

	/* logOp(add, 100, 200)
	logOp(subtract, 100, 200) */

	add := logOp(add)
	subtract := logOp(subtract)
	add(100, 200)
	subtract(100, 200)
}

func logOp(oper func(int, int) int) func(int, int) {
	return func(x, y int) {
		fmt.Println("Before invocation")
		result := oper(x, y)
		fmt.Println("result = ", result)
		fmt.Println("After invocation")
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
