package main

import (
	"flag"
	"fmt"
	"strings"
)

func maskCard(card string) string {
	return card[:4] + " **** **** " + card[12:]
}

func normalizeCard(card string) string {
	card = strings.ReplaceAll(card, " ", "")
	card = strings.ReplaceAll(card, "-", "")
	return card
}

func main() {
	card := flag.String("card", "", "номер карты")
	flag.Parse()

	if *card == "" {
		fmt.Println("Ошибка: укажите номер карты через -card")
		return
	}

	normalized := normalizeCard(*card)

	if len(normalized) != 16 {
		fmt.Println("Ошибка: неверный формат карты")
		return
	}

	for _, ch := range normalized {
		if ch < '0' || ch > '9' {
			fmt.Println("Ошибка: номер должен содержать только цифры")
			return
		}
	}

	fmt.Println("Нормализованный номер:", normalized)
	fmt.Println("Длина:", len(normalized))
	fmt.Println("Карта:", maskCard(normalized))
}
