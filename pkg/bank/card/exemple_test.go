package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExemplePaymentSources() {
	fmt.Println(PaymentSources( []types.Card{
		{
			Balance: 1_000_00,
			PAN: "5555 6666 7777 8878",
			Active: true,
			Currency: "TJS",
 		},
	}))
	
fmt.Println(PaymentSources( []types.Card{
	{
		Balance: 1_000_00,
		PAN: "5678 9876 8765 8798",
		Active: false,
		Currency: "TJS",
	},
}))

	fmt.Println(PaymentSources( []types.Card{
		{
			Balance: -1_000_00,
			PAN: "9074 9934 8645 7642",
			Active: true,
			Currency: "TJS",
		},
	}))
   //Output: 
   //5555 6666 7777 8878(100000 TJS)
}