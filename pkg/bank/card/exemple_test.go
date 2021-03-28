package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExemplePaymentSources() {
	payments := [] types.Card{
		{
			Balance: 1_000_00,
			PAN: "01",
			Active: true,
		},
		{ 
			Balance: 1_000_00,
			PAN: "02",
			Active: false,
		},
		{
			Balance: -1_000_00,
			PAN: "03",
		    Active: true,
		},
		{
			Balance: 2_000_00,
			PAN: "04",
		    Active: true,
		},
    } 
	fmt.Println(PaymentSources(payments))
	//Output:
	//01
	//04
}