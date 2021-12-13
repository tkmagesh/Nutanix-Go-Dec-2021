package main

import "fmt"

/* var s string */

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	/* type inference */
	/*
		var msg = "Hello World!"
	*/

	/* := can be used only in a function */
	msg := "Hello World!"
	fmt.Println(msg)

	/* Declaring multiple variables */
	/*
		var x int
		var y int
		var result int
		var message string

		x = 100
		y = 200
		result = x + y
		message = "Result ="
	*/

	/*
		var x int = 100
		var y int = 200
		var result int = x + y
		var message = "Result ="
	*/

	/*
		var x, y, result int
		var message string
		x = 100
		y = 200
		result = x + y
		message = "Result ="
	*/

	/*
		var x, y int = 100, 200
		var result int = x + y
		var message string = "Result ="
	*/

	/*
		var x, y = 100, 200
		var result = x + y
		var message = "Result ="
	*/

	/*
		var x, y, message = 100, 200, "Result ="
		var result = x + y
		fmt.Printf("Output : %s %d\n", message, result)
	*/

	/*
		var (
			x, y, message = 100, 200, "Result ="
			result        = x + y
		)
	*/
	/*
		var (
			x, y    int    = 100, 200
			message string = "Result ="
			result  int    = x + y
		)
	*/

	x, y, message := 100, 200, "Result ="
	result := x + y
	fmt.Printf("Output : %s %d\n", message, result)

	var s string
	if false {
		fmt.Println(s)
	}

	const no int = 100

	/* const (
		red   = iota
		green = iota
		blue  = iota
	) */

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	const (
		red = iota + 5
		green
		_
		_
		_
		blue
	)
	const (
		pen = iota
		pencil
		marker
	)
	fmt.Printf("red = %d, green = %d, blue = %d\n", red, green, blue)
	fmt.Printf("pen = %d, pencil = %d, marker = %d\n", pen, pencil, marker)
}
