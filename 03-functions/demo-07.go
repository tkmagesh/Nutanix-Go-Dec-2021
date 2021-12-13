package main

import (
	"errors"
	"fmt"
)

var (
	InvalidArguments = errors.New("invalid arguments")
)

func main() {
	q, r, err := divide(100, 0)
	if err == nil {
		fmt.Println("something went wrong -> ", err)
		return
	}
	fmt.Println(q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, errors.New("invalid arguments")
	}
	return x / y, x % y, nil
}
*/

func divide(x, y int) (quotient int, remainder int, err error) {
	if y == 0 {
		err = InvalidArguments
		return
	}
	quotient, remainder = x/y, x%y
	return
}
