package main

import (
	"fmt"
)

func main() {
	for {
		var amount float64
		var input int

		fmt.Print("Введите сумму перевода: ")
		fmt.Scan(&amount)

		if amount < 500 {
			fmt.Println("Введите сумму больше 500 сум!")
			continue
		}
		if amount > 15_000_000 {
			fmt.Println("Введите сумму меньше 15.000.000 сум!")
			continue
		}

		fmt.Print("Alif карта? (1-да/0-нет): ")
		fmt.Scan(&input)

		if input == 1 {
			fmt.Println("Ты выбрал ДА")
			break
		} else if input == 0 {
			fmt.Println("Возврат в главное меню...")
			continue
		} else {
			fmt.Println("Некорректный ввод")
			continue
		}
	}
}
