package main

import "fmt"

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}

	/*
		split the data into 2 equal sets
		aggregate the data sets using the sum (should happen concurrently)
		collect the intermediate result
		aggregate the intermediate result
		print the final result
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)
	dataset1 := data[:len(data)/2]
	dataset2 := data[len(data)/2:]

	go func(ch chan int, data []int) {
		ch <- sum(data...)
	}(ch1, dataset1)

	go func(ch chan int, data []int) {
		ch <- sum(data...)
	}(ch2, dataset2)

	result := <-ch1 + <-ch2
	fmt.Println(result)
}

func sum(nos ...int) int {
	total := 0
	for _, v := range nos {
		total += v
	}
	return total
}
