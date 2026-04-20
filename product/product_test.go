package product

import "fmt"

func ExampleMonthlyPayment() {
	p := Product{
		ID:     1,
		Name:   "Xiaomi Redmi Note 14 6/128 GB",
		Price:  2499000,
		Months: 12,
	}

	fmt.Println(monthlyPayment(p))
	// Output:
	// 208250
}
