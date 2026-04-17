package main

import "fmt"

func main() {
	product := "iPhone 15 Pro"
	brand := "Apple"
	available := true
	amount := int64(12_990_000)

	plan := amount / 12

	fmt.Printf("===== Alifshop =====\n")
	fmt.Printf("Товар:     %s\n", product)
	fmt.Printf("Бренд:     %s\n", brand)
	fmt.Printf("Цена:      %d сум\n", amount)
	fmt.Printf("В наличии: %t\n", available)
	fmt.Printf("Рассрочка: 12 мес → %d сум/мес\n", plan)
	fmt.Println("====================")
}
