package main

import (
	"fmt"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("application panicked, err = ", err)
		}
	}()
	result, e := divideClient(100, 0)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("result = ", result)
}

func divideClient(x, y int) (result int, err error) {
	/* defer func() {
		if e := recover(); e != nil {
			err = errors.New("divide error ")
			return
		}
	}() */
	result = divide(x, y)
	return
}

func divide(x, y int) (result int) {
	if y == 0 {
		panic("Cannot divide by zero")
	}
	return x / y
}
