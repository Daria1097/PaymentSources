package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExemplePaymentSources() {
	payments := [] types.Card{
		{
			Balance: 1_000_00,
			Active: true,
		},
		{ 
			Balance: 1_000_00,
			Active: false,
		},
		{
			Balance: -1_000_00,
		    Active: true,
		},
    }

	sources:= PaymentSources(payments)
	fmt.Println(sources)
	//Output: 2
}