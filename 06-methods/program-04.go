package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

/*
func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}
*/

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

	products.AddProduct(Product{106, "Eraser", 10, 50, "Stationary"})
	fmt.Println(products)
}

/*
func (products Products) FormatProducts() string {
	var result string
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p.Format())
	}
	return result
}
*/
//implementation of the fmt.Stringer interface
func (products Products) String() string {
	var result string
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p)
	}
	return result
}

func (products *Products) AddProduct(product Product) {
	*products = append(*products, product)
}
