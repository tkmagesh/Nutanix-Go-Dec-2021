package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40, 50))
}

func sum(firstNo int, nos ...int) int {
	//nos => list of int (slice)
	//len(nos) => number of values in nos
	//nos[0] => first value in the nos list

	result := firstNo
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}
