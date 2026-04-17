package main

import "fmt"

func main() {

	sender := "Sherali"
	receiver := "Alisher"
	amount := int64(500_000)
	comission := int64(5_000)
	total := amount + comission

	fmt.Println("========= Чек =========")
	fmt.Printf("Отправитель: %s\nПолучатель: %s\nСумма: %d\nКомиссия: %d\nИтого: %d\n", sender, receiver, amount, comission, total)
	fmt.Println("=======================")
}
