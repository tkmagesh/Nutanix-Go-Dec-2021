package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no
	//derefernecing ptr -> value
	var z = *noPtr
	fmt.Println(noPtr, z)

	increment(noPtr)
	fmt.Println(no)

	var n1, n2 int = 10, 20
	swap(&n1, &n2)
	fmt.Println(n1, n2)
}

func increment(x *int) {
	(*x)++
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
