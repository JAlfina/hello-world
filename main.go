package main

import (
	"fmt"
	"hello-world-1/internal"
)

func main() {
	var senderFirstName string
	var senderLastName string
	var receiverFirstName string
	var receiverLastName string
	var cardNumber string
	var amount int
	var input int

	fmt.Print("Имя отправителя: ")
	fmt.Scan(&senderFirstName)

	fmt.Print("Фамилия отправителя: ")
	fmt.Scan(&senderLastName)

	fmt.Print("Имя получателя: ")
	fmt.Scan(&receiverFirstName)

	fmt.Print("Фамилия получателя: ")
	fmt.Scan(&receiverLastName)

	fmt.Print("Номер карты получателя (16 цифр): ")
	fmt.Scan(&cardNumber)

	if len(cardNumber) != 16 {
		fmt.Println("Ошибка: номер карты должен содержать 16 цифр")
		return
	}

	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	if amount < 500 {
		fmt.Println("Введите сумму больше 500 сум!")
		return
	}

	if amount > 15000000 {
		fmt.Println("Введите сумму меньше 15 000 000 сум!")
		return
	}

	fmt.Print("Alif карта? (1-да/0-нет): ")
	fmt.Scan(&input)

	if input != 0 && input != 1 {
		fmt.Println("Некорректный ввод")
		return
	}

	isAlif := input == 1

	receipt := internal.BuildReceipt(
		senderFirstName,
		senderLastName,
		receiverFirstName,
		receiverLastName,
		cardNumber,
		amount,
		isAlif,
	)

	fmt.Println()
	fmt.Println(internal.FormatReceipt(receipt))
}
