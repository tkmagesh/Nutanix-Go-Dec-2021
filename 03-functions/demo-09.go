package main

import (
	"errors"
	"fmt"
)

var CannotDivideByZero = errors.New("cannot divide by zero")

func main() {

	defer func() {
		err := recover()
		if err == CannotDivideByZero {
			//do something
			return
		}
		if err != nil {
			fmt.Printf("type of err = %T\n", err)
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
		panic(CannotDivideByZero)
	}
	return x / y
}
