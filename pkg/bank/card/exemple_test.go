package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExemplePaymentSources() {
	cards := []types.Card{
	{
		Balance: 1_000_00,
		Active: true,
 	},
	{
		Balance: 3_000_00,
		Active: true,
	},
	{
		Balance: -1_000_00,
		Active: true,
	},
	{
		Balance: 2_000_00,
		Active: true,
	},
	{
		Balance: 10_000_00,
		Active: false,
	},
  }
   fmt.Println(PaymentSources(cards))
   //Output: 
   //100000
   //300000
   //200000
}