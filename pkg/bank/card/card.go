package card

import (
	"github.com/bakhtiyour/bank/pkg/bank/types"
	//"fmt"
)



func IssueCard(currency types.Currency, color string, name string) types.Card {
card :=types.Card {
ID:			1000,
PAN:		"5058 xxxx xxxx 0001",
Balance:	0,
Currency:	currency,
Active:		true,
Color:		color,
Name:		name,
}

return card
}

func Withdraw (card types.Card, amount types.Money) types.Card {
if card.Balance<amount{
	return card
}

if amount<0{
	return card
}

if amount>20000_00{
	return card
}

if card.Active==false{
	return card
}
card.Balance=card.Balance-amount	
return card
}


func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}

	if amount > 50_000_00 {
		return
	}
	if amount<0{
		return
	}

	card.Balance +=amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	
	bonus:=int(card.MinBalance)*percent*daysInMonth/(daysInYear*100)
	if !card.Active {
		return
	}

	if card.Balance < 0 {
		return
	}
	if bonus>500000{
		//card.Balance+=500000
		return
	}

	card.Balance +=types.Money(bonus)
}

func Total(cards []types.Card) types.Money {
	total:=types.Money(0)
	
	for i:=0; i<len(cards); i++{
		cardBalance:=cards[i].Balance
		if !cards[i].Active{
	 		cardBalance=0
		}

		if cardBalance<0{
			cardBalance=0
		}

	total +=cardBalance
	
	}
	return total
}

func PaymentSources(cards []types.Card) []types.PaymentSource{
	var paySource []types.PaymentSource
	
	for i:=0;i<len(cards);i++{
		if cards[i].Balance>0 {
			if cards[i].Active{
				payAppend:=types.PaymentSource{Type: "card", Number: string(cards[i].PAN), Balance: cards[i].Balance, }
				paySource=append(paySource, payAppend)
			}
			
		}

	}
	return paySource
}

