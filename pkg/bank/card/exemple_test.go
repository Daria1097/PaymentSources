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
		Active: true,
	},
	{
		Balance: 2_000_00,
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
   //01
   //04
   
}