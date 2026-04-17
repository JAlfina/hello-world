package main

import "fmt"

func main() {
	product := "iPhone 15 Pro"
	brand := "Apple"
	available := true
	amount := int64(12_990_000)

	plan := amount / 12

	fmt.Printf("===== Alifshop =====\nТовар: %s\nБренд: %s\nЦена: %d сум\nВ наличии: %t\nРассрочка: 12 мес → %d сум/мес\n", product, brand, amount, available, plan)
	fmt.Print("====================")
}
