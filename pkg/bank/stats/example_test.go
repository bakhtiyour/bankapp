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