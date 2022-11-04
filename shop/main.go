package main

import (
	"fmt"
	"shop/admin"
)

var (
	bank = admin.Bank{
		Name:  "bank",
		Cards: []admin.Card{},
	}
	card1 = admin.Card{
		Name:    "crd-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = admin.Card{
		Name:    "crd-2",
		Balance: 5,
		Bank:    &bank,
	}
	user = admin.User{
		Name: "user-1",
		Card: &card1,
	}
	user2 = admin.User{
		Name: "user-2",
		Card: &card2,
	}
	prod = admin.Product{
		Name:  "Cheese",
		Price: 150,
	}
	shop = admin.Shop{
		Name: "Duken",
		Products: []admin.Product{
			prod,
		},
	}
)

func main() {
	println("[Bank] Vypusk card")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]", user.Name)
	err := shop.Sell(user, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("[%s]", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}

}
