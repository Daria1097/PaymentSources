package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExemplePaymentSources() {
	fmt.Println(PaymentSources( []types.Card{
		{
			Balance: 1_000_00,
			Active: true,
 		},
	}))
	
fmt.Println(PaymentSources( []types.Card{
	{
		Balance: 1_000_00,
		Active: false,
	},
}))

	fmt.Println(PaymentSources( []types.Card{
		{
			Balance: -1_000_00,
			Active: true,
		},
	}))
   //Output: 
   //5555 6666 7777 8878(100000 TJS)
}