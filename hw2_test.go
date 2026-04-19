package main

import "fmt"

func Example_maskCard() {
	fmt.Println(maskCard("1234567812345678"))
	// Output:
	// 1234 **** **** 5678
}

func Example_normalizeCard_spaces() {
	fmt.Println(normalizeCard("1234 5678 1234 5678"))
	// Output:
	// 1234567812345678
}

func Example_normalizeCard_dashes() {
	fmt.Println(normalizeCard("1234-5678-1234-5678"))
	// Output:
	// 1234567812345678
}

func Example_maskCard_panic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic")
		}
	}()

	fmt.Println(maskCard("123123123"))

	// Output:
	//
}
