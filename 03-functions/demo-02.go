package main

import "fmt"

func main() {
	/*
		sayHello := func() {
			fmt.Println("Hello World!")
		}
	*/
	var sayHello func()
	sayHello = func() {
		fmt.Println("Hello World!")
	}
	sayHello()
	fmt.Printf("type of sayHello = %T\n", sayHello)

	var oper func(int, int) int
	oper = func(x, y int) int {
		return x + y
	}
	fmt.Println(oper(100, 200))

	/* anonymous functions */
	func() {
		fmt.Println("An anonymous function")
	}()

	func(x, y int) {
		fmt.Println("Adding in an anonymous function, result =", x+y)
	}(100, 200)

	result := func(x, y int) int {
		fmt.Println("Adding in an anonymous function returning result")
		return x + y
	}(100, 200)
	fmt.Println("result =", result)
}

/*
func sayHello() {
	fmt.Println("Hello World!")
}
*/
