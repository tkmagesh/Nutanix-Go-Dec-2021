package main

import (
	"fmt"
	"sort"
)

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

func (products Products) SortByName() {
	sort.Sort(ByName{products})
}

func (products Products) SortByCost() {
	sort.Sort(ByCost{products})
}

func (products Products) SortByUnits() {
	sort.Sort(ByUnits{products})
}

func (products Products) SortByCategory() {
	sort.Sort(ByCategory{products})
}

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
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Sorting [default by id]")
	sort.Sort(products)
	fmt.Println(products)

	fmt.Println("Sorting [by name]")
	sort.Sort(ByName{products})
	fmt.Println(products)

	fmt.Println("Sorting [by cost]")
	products.SortByCost()
	fmt.Println(products)

	fmt.Println("Sorting [by units]")
	products.SortByUnits()
	fmt.Println(products)

	fmt.Println("Sorting [by category]")
	products.SortByCategory()
	fmt.Println(products)
}

//implementation of the fmt.Stringer interface
func (products Products) String() string {
	var result string
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p)
	}
	return result
}

//Create the APIs for sorting the products (USE sort.Sort())
// by id, name, cost, units, category

/* Interface implementation */
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

//to sort by name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//to sort by cost
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

//to sort by units
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

//to sort by category
type ByCategory struct {
	Products
}

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}
