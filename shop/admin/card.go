package admin

import "time"

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (card Card) CheckBalance() error {
	println("[Card] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)
	return card.Bank.CheckBalance(card.Name)
}