package main

import "fmt"

func main() {
	//Array
	fmt.Println("Array")
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 2, 4, 5}
	fmt.Println(nos)

	//iteration using index
	fmt.Println("Iteration using index")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(idx, nos[idx])
	}

	fmt.Println("Iteration using range")
	//iteration using range constructor
	for idx, val := range nos {
		nos[idx] = val * 2
	}
	for idx, val := range nos {
		fmt.Println(idx, val)
	}

	//creating a copy of an array
	newNos := nos
	newNos[0] = 100
	fmt.Println(newNos, nos)

	//Slice
	fmt.Println("Slice")
	var products []string = []string{"Pen", "Pencil", "Marker"}
	fmt.Printf("Address of products = %p\n", products)
	fmt.Println(products)
	fmt.Println("Capacity of products (before adding) = ", cap(products))

	//adding new values to the slice
	products = append(products, "ScribblePad")
	fmt.Printf("Address of products = %p\n", products)
	fmt.Println(products)
	fmt.Println("Capacity of products (after adding) = ", cap(products))

	products = append(products, "stylus")
	fmt.Printf("Address of products = %p\n", products)
	fmt.Println(products)
	fmt.Println("Capacity of products (after adding) = ", cap(products))

	/*
		products = append(products, "p1", "p2", "p3")
		fmt.Println("Capacity of products (after adding) = ", cap(products))

		dupProducts := append(products, "P1001")
		dupProducts[0] = "Dummy-Product"
		fmt.Println(products, dupProducts)

		products = append(products, "another-dummy-product")
		fmt.Println(products, dupProducts)
	*/
	//slicing
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to end
		[ : hi] => from 0 to hi-1
		[lo : lo] => empty slice
		[lo : lo+1] => single element slice with [lo]
		[:] => copy of the slice
	*/
	fmt.Println(products)
	fmt.Println("products[1:3] => ", products[1:4])

	//Map
	fmt.Println("Maps")
	/*
		var productRanks map[string]int
		productRanks["pen"] = 1
		productRanks["pencil"] = 5
	*/
	var productRanks map[string]int = map[string]int{
		"pen":         1,
		"pencil":      4,
		"marker":      3,
		"scribblepad": 5,
	}
	fmt.Println(productRanks)

	fmt.Println("Iterating a map")
	for key, value := range productRanks {
		fmt.Println(key, value)
	}

	fmt.Println("adding a new key/value pair")
	productRanks["stylus"] = 2
	fmt.Println(productRanks)

	fmt.Println("Checking if a key exists")
	p := "stapler"
	if val, exists := productRanks[p]; exists {
		fmt.Printf("rank of %s is %d\n", p, val)
	} else {
		fmt.Printf("[%s] doesnot exist, but val = %d\n", p, val)
	}

	fmt.Println("Removing a key/value pair")
	delete(productRanks, "pen")
	fmt.Println(productRanks)

	//make function
	numbers := make([]int, 5, 100)
	numbers[0] = 100
	numbers[1] = 200
	numbers = append(numbers, 10, 20, 30)
	fmt.Println("numbers = ", numbers)

}
