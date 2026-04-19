package product

import (
	"strconv"
	"strings"
)

type Product struct {
	Name    string
	Brand   string
	Price   int // в тийинах
	InStock bool
}

const template = `===== Alifshop =====
Товар:    {name}
Бренд:    {brand}
Цена:     {price} сум
В наличии: {inStock}
Рассрочка: 12 мес → {firstAmount} сум/мес
====================`

func Converter(product Product, firstAmount int) string {
	// перевод из тийинов в сумы
	priceSum := product.Price / 100
	firstAmountSum := firstAmount / 100

	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", strconv.Itoa(priceSum))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", strconv.Itoa(firstAmountSum))

	return text
}
