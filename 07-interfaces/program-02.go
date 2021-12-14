package main

import "fmt"

type MyType struct {
	i int
}

func (myType MyType) String() string {
	return "My Type"
}

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

	x = MyType{i: 10}
	//x = struct{ i int }{10}

	switch val := x.(type) {
	case int:
		fmt.Println("Double of x is ", val*2)
	case string:
		fmt.Println("Length of x is ", len(val))
	case bool:
		fmt.Println("x is a boolean")
	case struct{ i int }:
		fmt.Println("x is a struct")
	case MyType:
		fmt.Println("x is a MyType")
	case []int:
		fmt.Println("x is a slice of integers")
	case map[string]int:
		fmt.Println("x is a map of strings to integers")
	case fmt.Stringer:
		fmt.Println("x is a fmt.Stringer")
	default:
		fmt.Println("Unknown type")
	}
}
