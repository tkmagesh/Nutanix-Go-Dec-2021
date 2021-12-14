package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

//implementation of the fmt.Stringer interface
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

type Products []Product

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	//fmt.Println(products.FormatProducts())
	fmt.Println(products)

	/
}

//implementation of the fmt.Stringer interface
func (products Products) String() string {
	var result string
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p)
	}
	return result
}

//Create the APIs for sorting the products 
	// by id, name, cost, units, category