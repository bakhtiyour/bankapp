package payment

import (
	"bank/pkg/bank/types"
	//"fmt"
)
func Max(payments []types.Payment) types.Payment{
	maxPayment:=payments[0]
	
	for i:=0; i<len(payments); i++{
		if payments[i].Amount>maxPayment.Amount{
			maxPayment=payments[i]
			
		}
	}
	return maxPayment
}