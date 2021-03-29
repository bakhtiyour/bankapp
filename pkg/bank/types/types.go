package types

//Money представляет собой денежную сумму в минимальных единицах (дирамы, центы и т.д.)
type Money int64

//Currency представляет код валюты
type Currency string

//Коды валют
const(
	TJS Currency ="TJS"
	RUB Currency ="RUB"
	USD Currency ="USD"
)

//PAN представляет номер карты
type PAN string

//Payment представляет информацию о платеже
type Payment struct {
	ID int
	Amount Money
	Category Category
}
//Card представляет информацию о платёжной карте
type Card struct {
	ID			int
	PAN 		PAN
	Balance		Money
	Currency 	Currency
	Color 		string
	Name 		string
	Active		bool
	MinBalance	Money
}

type PaymentSource struct{
	Type string
	Number string
	Balance Money
}

type Category string


