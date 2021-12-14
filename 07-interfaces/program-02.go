package main

import "fmt"

func main() {
	var x interface{}
	x = 10
	/*
		x = "Hello"
		x = true
		x = struct{}{}
	*/
	if val, ok := x.(int); ok {
		fmt.Println("x is an integer")
		fmt.Printf("x + 10 = %d\n", val+10)
	} else {
		fmt.Println("x is not an integer")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("Double of x is ", val*2)
	case string:
		fmt.Println("Length of x is ", len(val))
	case bool:
		fmt.Println("x is a boolean")
	case struct{}:
		fmt.Println("x is a struct")
	case []int:
		fmt.Println("x is a slice of integers")
	case map[string]int:
		fmt.Println("x is a map of strings to integers")
	default:
		fmt.Println("Unknown type")
	}
}
