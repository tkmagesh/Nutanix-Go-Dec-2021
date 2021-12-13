package main

import "fmt"

func main() {
	sayHello()
	fmt.Println("isPrime(97) = ", isPrime(97))
	//fmt.Println(divide(100, 7))
	/*
		q, r := divide(100, 7)
		fmt.Printf("quotient = %d, remainder = %d\n", q, r)
	*/
	q, _ := divide(100, 7)
	fmt.Printf("Quotient = %d\n", q)
}

func sayHello() {
	fmt.Println("Hello World!")
}

/*
func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
*/

func isPrime(no int) (result bool) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			result = false
			return
		}
	}
	result = true
	return result
}

//functions returning more than one result
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
