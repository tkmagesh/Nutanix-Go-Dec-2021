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
	product Product
	expiry  string
}

func NewPerishableProduct(id int, name string, cost float64, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{
		product: Product{id, name, cost, units, category},
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
	fmt.Println(pp.product.Name, pp.product.Cost)

	//fmt.Println(ToString(pp))
	fmt.Println(ToString(pp.product))

	//applyDiscount(pp)
	applyDiscount(&pp.product, 10)
	fmt.Println(ToString(pp.product))
}

func ToString(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func applyDiscount(p *Product, discount float64) {
	//(*p).Cost = (*p).Cost * ((100 - discount) / 100)
	p.Cost = p.Cost * ((100 - discount) / 100)
}
