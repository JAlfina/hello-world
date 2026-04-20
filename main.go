package main

import "hello-world-1/product"

func main() {
	p := product.Product{
		ID:     1,
		Name:   "Xiaomi Redmi Note 14 6/128 GB",
		Price:  2499000,
		Months: 12,
	}

	product.PrintProductForMain(p)
}
