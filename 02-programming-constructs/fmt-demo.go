package main

import "fmt"

func main() {
	var no int
	fmt.Println("Enter a number :")
	//fmt.Scanln(&no)
	fmt.Scanf("%d\n", &no)
	fmt.Println("no = ", no)
}
