package stats

import (

	"github.com/bakhtiyour/bank/pkg/bank/types"
)

func Avg(payments []types.Payment)types.Money {
	amountSum :=types.Money(0)
	amountCount:=0
	for i:=0; i<len(payments); i++{
		amountSum +=payments[i].Amount
		amountCount+=1
	}
	avgPayment:=amountSum/types.Money(amountCount)
	return avgPayment
}