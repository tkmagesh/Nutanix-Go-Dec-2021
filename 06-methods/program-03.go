/* Association */
package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	expiry string
}

func NewPerishableProduct(id int, name string, cost float64, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{id, name, cost, units, category},
		expiry:  expiry,
	}
}

func main() {
	/*
		pp := PerishableProduct{
			product: Product{200, "Corn", 5, 100, "Food"},
			expiry:  "5 Days",
		}
	*/

	pp := NewPerishableProduct(200, "Corn", 5, 100, "Food", "5 Days")

	//fmt.Println(pp.Product.Name, pp.Product.Cost)
	fmt.Println(pp.Name, pp.Cost)

	//fmt.Println(ToString(pp))
	//fmt.Println(ToString(pp.Product))
	fmt.Println(pp.ToString())

	//applyDiscount(pp)
	//applyDiscount(&pp.Product, 10)
	pp.applyDiscount(10)
	//fmt.Println(ToString(pp.Product))
	fmt.Println(pp.ToString())
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (p *Product) applyDiscount(discount float64) {
	//(*p).Cost = (*p).Cost * ((100 - discount) / 100)
	p.Cost = p.Cost * ((100 - discount) / 100)
}

//Overriding the ToString method
func (pp PerishableProduct) ToString() string {
	return fmt.Sprintf("%s, expiry = %s", pp.Product.ToString(), pp.expiry)
}
