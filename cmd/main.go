package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	
	newCards:=[]types.Card{
	{
		Balance: 10000,
		PAN: "5051",
		Active: false,
	},
	{
		Balance: -20000,
		PAN: "5052",
		Active: true,
	},
	{
		Balance: 20000,
		PAN: "5053",
		Active: true,
	},
}
	tot:=card.PaymentSources(newCards)
	fmt.Println(tot)
}