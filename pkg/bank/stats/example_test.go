package stats
import(
	"fmt"
	"github.com/bakhtiyour/bank/pkg/bank/types"
	
)
func ExampleAvg() {
	newPayments:=[]types.Payment{
		{
			Amount: 10000,
		},
		{
			Amount: 10000,
		},
		{
			Amount: 10000,
		},
	} 
	result:=Avg(newPayments)
	fmt.Println(result)
	// Output: 10000
}

func ExampleTotalInCategory() {
	newPayments:=[]types.Payment{
		{
			Amount: 10000,
			Category: "mobile",
		},
		{
			Amount: 10000,
			Category: "mobile",
		},
		{
			Amount: 10000,
			Category: "bank",
		},
	} 
	result:=TotalInCategory(newPayments, "mobile")
	fmt.Println(result)
	// Output: 20000
}