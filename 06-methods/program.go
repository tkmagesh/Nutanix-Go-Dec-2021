package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	//var p Product
	//var p Product = Product{}
	//var p Product = Product{101, "Pen", 10, 100, "Stationary"}
	//var p Product = Product{Id: 101, Name: "Pen", Cost: 10}

	var p Product = Product{
		Id:       101,
		Name:     "Pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}
	//p := Product{}
	//fmt.Printf("%#v\n", p)
	//fmt.Println(ToString(p))
	fmt.Println(p.ToString())
	fmt.Println("Applying 10% discount")
	//applyDiscount(&p, 10)
	p.applyDiscount(10)
	fmt.Println(p.ToString())
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (p *Product) applyDiscount(discount float64) {
	//(*p).Cost = (*p).Cost * ((100 - discount) / 100)
	p.Cost = p.Cost * ((100 - discount) / 100)
}
