package card
import(
	"fmt"
	"github.com/bakhtiyour/bank/pkg/bank/types"
	
)
func ExampleWithdraw_positive() {
	result:=Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	result:=Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 0
}

func ExampleWithdraw_inactive() {
	result:=Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

func ExampleWithdraw_limit() {
	result:=Withdraw(types.Card{Balance: 30_000_00, Active: false}, 21_000_00)
	fmt.Println(result.Balance)
	// Output: 3000000
}

func ExampleDeposit_positive() {
	card:=types.Card{Balance: -30_000_00, Active: true}
	Deposit(&card, 40_000_00)
	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleDeposit_inactive() {
	card:=types.Card{Balance: -30_000_00, Active: false}
	Deposit(&card, 30_000_00)
	fmt.Println(card.Balance)
	// Output: -3000000
}

func ExampleDeposit_limit() {
	card:=types.Card{Balance: -30_000_00, Active: true}
	Deposit(&card, 60_000_00)
	fmt.Println(card.Balance)
	// Output: -3000000
}

func ExampleAddBonus_positive() {
	card:=types.Card{Balance: 30_000_00, Active: true, MinBalance: 1000000}
	AddBonus(&card, 3,30,365)
	fmt.Println(card.Balance)
	// Output: 3002465
}

func ExampleTotal() {
	newCards:=[]types.Card{
		{
			Balance: 10000,
			Active: false,
		},
		{
			Balance: -20000,
			Active: true,
		},
		{
			Balance: 20000,
			Active: true,
		},
	}

	fmt.Println(Total(newCards))
	
	//Output: 20000
	
}