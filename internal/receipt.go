package internal

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Receipt struct {
	SenderFirstName   string
	SenderLastName    string
	ReceiverFirstName string
	ReceiverLastName  string
	CardNumber        string
	Amount            int
	Commission        int
	Total             int
	TransactionID     int
	DateTime          string
	MaskedCard        string
	Status            string
}

func GenerateTransactionID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000000) + 100000000
}

func FormatDateTime() string {
	return time.Now().Format("02.01.2006 15:04")
}

func MaskCard(cardNumber string) string {
	if len(cardNumber) < 4 {
		return "UZCARD**0000"
	}
	last4 := cardNumber[len(cardNumber)-4:]
	return "UZCARD**" + last4
}

func ToUpperFullName(firstName, lastName string) string {
	fullName := lastName + " " + firstName
	return strings.ToUpper(fullName)
}

func BuildReceipt(
	senderFirstName, senderLastName string,
	receiverFirstName, receiverLastName string,
	cardNumber string,
	amount int,
	isAlif bool,
) Receipt {
	commission := CalculateCommission(amount, isAlif)
	total := amount + commission

	return Receipt{
		SenderFirstName:   senderFirstName,
		SenderLastName:    senderLastName,
		ReceiverFirstName: receiverFirstName,
		ReceiverLastName:  receiverLastName,
		CardNumber:        cardNumber,
		Amount:            amount,
		Commission:        commission,
		Total:             total,
		TransactionID:     GenerateTransactionID(),
		DateTime:          FormatDateTime(),
		MaskedCard:        MaskCard(cardNumber),
		Status:            "Исполнено",
	}
}

func CalculateCommission(amount int, isAlif bool) int {
	if isAlif {
		return 0
	}
	return amount * 5 / 1000
}

func FormatReceipt(r Receipt) string {
	return fmt.Sprintf(
		`========= Электронный чек =========
Отправитель:      %s
Получатель:       %s
Номер транзакции: %d
Счёт зачисления:  %s
Дата и время:     %s
Сумма:            %d сум
Комиссия:         %d сум
Итого:            %d сум
AO ALIF TECH · Лицензия ЦБ РУз № 000010
Статус:           %s`,
		ToUpperFullName(r.SenderFirstName, r.SenderLastName),
		ToUpperFullName(r.ReceiverFirstName, r.ReceiverLastName),
		r.TransactionID,
		r.MaskedCard,
		r.DateTime,
		r.Amount,
		r.Commission,
		r.Total,
		r.Status,
	)
}
