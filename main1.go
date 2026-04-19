package main

import (
	"fmt"
	"strings"

	"hello-world-1/internal/commission"
)

func main() {
	var amount int
	var input int
	var fullName string

	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	if !commission.CheckAmount(amount) {
		fmt.Println("Ошибка: сумма должна быть от 500 до 15 000 000 сум")
		return
	}

	fmt.Print("Alif карта? (1-да/0-нет): ")
	fmt.Scan(&input)

	if input != 0 && input != 1 {
		fmt.Println("Некорректный ввод")
		return
	}

	isAlif := input == 1
	fee := commission.CalculateCommission(amount, isAlif)
	total := commission.CalculateTotal(amount, fee)

	fullName = "Shokirov Shuhrat"

	fmt.Println()
	fmt.Println("========= Чек =========")
	fmt.Println()
	fmt.Println("Услуга:", strings.ToUpper(fullName))
	fmt.Println()
	fmt.Printf("Сумма: %d сум\n", amount)
	fmt.Printf("Комиссия: %d сум\n", fee)
	fmt.Printf("Итого: %d сум\n", total)
	fmt.Println()
	fmt.Println("Статус: Исполнено")
	fmt.Println()
	fmt.Println("=======================")
}
