package product

import "fmt"

type ProductID int
type Price int
type Months int

type Product struct {
	ID     ProductID
	Name   string
	Price  Price
	Months Months
}

func monthlyPayment(p Product) int {
	if p.Months < 3 || p.Months > 24 {
		return 0
	}

	return int(p.Price / Price(p.Months))
}

func printProduct(p Product) {
	if p.Months < 3 || p.Months > 24 {
		fmt.Println("Ошибка: срок рассрочки должен быть от 3 до 24 месяцев")
		return
	}

	fmt.Println("===== Alifshop =====")
	fmt.Printf("ID:       %d\n", p.ID)
	fmt.Printf("Товар:    %s\n", p.Name)
	fmt.Printf("Цена:     %d сум\n", p.Price)
	fmt.Printf("Срок:     %d мес\n", p.Months)
	fmt.Printf("Платёж:   %d сум/мес\n", monthlyPayment(p))
	fmt.Println("====================")
}
