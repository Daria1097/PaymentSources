package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExemplePaymentSources() {
	sources:= []types.Card{
		{
			Balance: -10_000_00,
			Active: true,
		 },
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: false ,
		},
	}
	   fmt.Println(PaymentSources(sources))
	   //Output:
	   //01
	   //04
	}	